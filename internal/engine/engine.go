package engine

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/juju/errors"
	"github.com/temoto/vender/helpers"
	"github.com/temoto/vender/log2"
)

type ErrNotResolved struct{ msg string }

func NewErrNotResolved(action string) ErrNotResolved {
	return ErrNotResolved{msg: fmt.Sprintf("action=%s not resolved", action)}
}
func (e ErrNotResolved) Error() string { return e.msg }

type Engine struct {
	Log     *log2.Log
	lk      sync.RWMutex
	actions map[string]Doer
}

func NewEngine(log *log2.Log) *Engine {
	self := &Engine{
		Log:     log,
		actions: make(map[string]Doer, 128),
	}
	self.actions["ignore(?)"] = FuncArg{
		Name: "ignore(?)",
		F:    func(context.Context, Arg) error { return nil }}
	self.actions["sleep(100ms)"] = Sleep{Duration: 100 * time.Millisecond}
	return self
}

func (self *Engine) Register(action string, d Doer) {
	self.lk.Lock()
	self.actions[action] = d
	self.lk.Unlock()
}

func (self *Engine) RegisterNewFunc(name string, fun func(context.Context) error) {
	self.Register(name, Func{
		Name: name,
		F:    fun,
	})
}

func (self *Engine) RegisterNewSeq(name string, ds ...Doer) {
	tx := NewSeq(name)
	for _, d := range ds {
		tx.Append(d)
	}
	self.Register(name, tx)
}

func (self *Engine) RegisterParse(name, scenario string) error {
	d, err := self.ParseText(name, scenario)
	if err != nil {
		err = errors.Annotatef(err, "engine.RegisterParse() name=%s scenario=%s", name, scenario)
		return err
	}
	self.Register(name, d)
	return nil
}

var reActionArg = regexp.MustCompile(`^(.+)\((\d+|\?)\)$`)

func (self *Engine) resolve(action string) (Doer, error) {
	// self.Log.Debugf("engine.resolve action=%s", action)
	self.lk.RLock()
	defer self.lk.RUnlock()
	d, ok := self.actions[action]
	if ok {
		// self.Log.Debugf("engine.resolve action=%s resolved d=%v", action, d)
		return d, nil
	}

	match := reActionArg.FindStringSubmatch(action)
	if match == nil {
		// self.Log.Debugf("engine.resolve action=%s match=nil", action)
		return nil, NewErrNotResolved(action)
	}

	normalized := match[1] + "(?)"
	argString := match[2]
	d, ok = self.actions[normalized]
	// self.Log.Debugf("engine.resolve action=%s match=%v normalized=%s ok=%t", action, match, normalized, ok)
	if !ok {
		self.Log.Debugf("resolve action=%s normalized=%s not found", action, normalized)
		err := NewErrNotResolved(normalized)
		err.msg = fmt.Sprintf(FmtErrContext, action) + err.msg
		return nil, err
	}
	if argString != "?" {
		argn, err := strconv.Atoi(argString)
		if err != nil {
			self.Log.Debugf("resolve action=%s err=%s", action, err)
			return nil, errors.Annotatef(err, FmtErrContext, action)
		}
		var applied bool
		d, applied, err = ArgApply(d, Arg(argn))
		if err != nil {
			self.Log.Debugf("resolve action=%s err=%s", action, err)
			return nil, errors.Annotatef(err, FmtErrContext, action)
		}
		if !applied {
			self.Log.Debugf("resolve action=%s arg=%v not applied", action, argString)
			err = ErrArgNotApplied
			return nil, errors.Annotatef(err, FmtErrContext, action)
		}
	}
	return d, nil
}

func (self *Engine) Resolve(action string) Doer {
	d, err := self.resolve(action)
	if err != nil {
		self.Log.Errorf("engine.Resolve action=%s err=%v", action, err)
		return Fail{E: err}
	}
	return d
}

func (self *Engine) List() []string {
	self.lk.RLock()
	r := make([]string, 0, len(self.actions))
	for k := range self.actions {
		r = append(r, k)
	}
	self.lk.RUnlock()
	return r
}

var reSleep = regexp.MustCompile(`sleep\((\d+m?s)\)`)

func (self *Engine) ResolveOrLazy(action string) (Doer, error) {
	self.lk.RLock()
	defer self.lk.RUnlock()
	d, ok := self.actions[action]
	if ok {
		// self.Log.Debugf("engine.ResolveOrLazy %s -> ok %#v", action, d)
		return d, nil
	}

	if m := reSleep.FindStringSubmatch(action); len(m) == 2 {
		duration, err := time.ParseDuration(m[1])
		if err != nil {
			return nil, errors.Trace(err)
		}
		return Sleep{duration}, nil
	}

	// self.Log.Debugf("engine.ResolveOrLazy %s -> lazy %#v", action, d)
	return &Lazy{Name: action, r: self.resolve}, nil
}

var reNotSpace = regexp.MustCompile(`\S+`)

func (self *Engine) ParseText(tag, text string) (Doer, error) {
	// TODO cache with github.com/hashicorp/golang-lru

	errs := make([]error, 0)
	words := reNotSpace.FindAllString(text, -1)

	tx := NewSeq(tag)
	for _, word := range words {
		d, err := self.ResolveOrLazy(word)
		if err != nil {
			return nil, errors.Annotatef(err, "scenario=%s unparsed=%s", text, word)
		}
		tx.Append(d)
	}
	return tx, helpers.FoldErrors(errs)
}

func (self *Engine) Exec(ctx context.Context, d Doer) error {
	err := d.Validate()
	if err == nil {
		err = d.Do(ctx)
	}
	return err
}

func (self *Engine) ExecList(ctx context.Context, tag string, list []string) []error {
	self.Log.Debugf("engine.ExecList tag=%s list=%v", tag, list)

	errs := make([]error, 0, len(list))
	for i, text := range list {
		itemTag := fmt.Sprintf("%s:%d", tag, i)
		d, err := self.ParseText(itemTag, text)
		if err == nil {
			err = self.Exec(ctx, d)
		}
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// Test `error` or `Doer` against ErrNotResolved
func IsNotResolved(x interface{}) bool {
	if x == nil {
		return false
	}
	e, _ := x.(error)
	if e == nil {
		if f, ok := x.(Fail); ok {
			e = f.E
		}
	}
	if e == nil {
		return false
	}
	e = errors.Cause(e)
	_, ok := e.(ErrNotResolved)
	return ok
}
