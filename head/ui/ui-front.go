package ui

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/juju/errors"
	"github.com/temoto/alive"
	"github.com/temoto/vender/currency"
	"github.com/temoto/vender/hardware/input"
	"github.com/temoto/vender/hardware/lcd"
	"github.com/temoto/vender/hardware/mdb/evend"
	"github.com/temoto/vender/head/money"
	"github.com/temoto/vender/head/tele"
)

const (
	DefaultCream = 4
	MaxCream     = 6
	DefaultSugar = 4
	MaxSugar     = 8
)

const modTuneTimeout = 3 * time.Second

var ScaleAlpha = []byte{
	0x94, // empty
	0x95,
	0x96,
	0x97, // full
	// '0', '1', '2', '3',
}

type UIMenuResult struct {
	Item  MenuItem
	Cream uint8
	Sugar uint8
}

func (self *UI) onFrontBegin(ctx context.Context) State {
	self.FrontResult = UIMenuResult{
		// TODO read config
		Cream: DefaultCream,
		Sugar: DefaultSugar,
	}

	// XXX FIXME custom business logic creeped into code TODO move to config
	if doCheckTempHot := self.g.Engine.Resolve("mdb.evend.valve_check_temp_hot"); doCheckTempHot != nil {
		err := doCheckTempHot.Validate()
		if errtemp, ok := err.(*evend.ErrWaterTemperature); ok {
			line1 := fmt.Sprintf(self.g.Config.UI.Front.MsgWaterTemp, errtemp.Current)
			self.display.SetLines(line1, self.g.Config.UI.Front.MsgWait)
			if e := self.wait(5 * time.Second); e.Kind == EventService {
				return StateServiceBegin
			}
			return StateFrontEnd
		} else if err != nil {
			self.g.Error(err)
			return StateBroken
		}
	}

	if err := self.g.Engine.ExecList(ctx, "on_front_begin", self.g.Config.Engine.OnFrontBegin); err != nil {
		self.g.Log.Errorf("on_front_begin err=%v", err)
		return StateBroken
	}

	var err error
	self.FrontMaxPrice, err = self.menu.MaxPrice(self.g.Log)
	if err != nil {
		self.g.Error(err)
		return StateBroken
	}
	self.g.Tele.State(tele.State_Nominal)
	return StateFrontSelect
}

func (self *UI) onFrontSelect(ctx context.Context) State {
	moneysys := money.GetGlobal(ctx)

	alive := alive.NewAlive()
	defer func() {
		alive.Stop() // stop pending AcceptCredit
		alive.Wait()
	}()
	go moneysys.AcceptCredit(ctx, self.FrontMaxPrice, alive.StopChan(), self.moneych)

	for {
	refresh:
		// step 1: refresh display
		credit := moneysys.Credit(ctx)
		if self.State == StateFrontTune { // XXX onFrontTune
			goto waitEvent
		}
		self.frontSelectShow(ctx, credit)

		// step 2: wait for input/timeout
	waitEvent:
		timeout := self.frontResetTimeout
		if self.State == StateFrontTune {
			timeout = modTuneTimeout
		}
		e := self.wait(timeout)
		switch e.Kind {
		case EventInput:
			switch e.Input.Key {
			case input.EvendKeyCreamLess, input.EvendKeyCreamMore, input.EvendKeySugarLess, input.EvendKeySugarMore:
				if next := self.onFrontTuneInput(e.Input); next != StateInvalid {
					return next
				}
				goto waitEvent // assume display was updated
			}
			if e.Input.IsDigit() {
				if next := self.onFrontSelectInput(ctx, e.Input); next != StateInvalid {
					return next
				}
				goto waitEvent // assume display was updated
			}
			switch self.State {
			case StateFrontSelect:
				if next := self.onFrontSelectInput(ctx, e.Input); next != StateInvalid {
					return next
				}
				goto waitEvent // assume display was updated

			case StateFrontTune:
				return StateFrontSelect

			default:
				panic(fmt.Sprintf("ui-front unhandled state=%s", self.State.String()))
			}

		case EventMoney:
			self.g.Log.Errorf("ui-front money event=%v", e.Money)
			switch e.Money.Name() {
			case money.EventAbort:
				self.g.Error(errors.Trace(moneysys.Abort(ctx)))
			}
			go moneysys.AcceptCredit(ctx, self.FrontMaxPrice, alive.StopChan(), self.moneych)
			goto refresh

		case EventService:
			return StateServiceBegin

		case EventTime:
			if self.State == StateFrontTune { // XXX onFrontTune
				return StateFrontSelect // "return to previous mode"
			}
			return StateFrontTimeout
		}
	}

	// external stop
	return StateFrontEnd
}

func (self *UI) frontSelectShow(ctx context.Context, credit currency.Amount) {
	config := self.g.Config.UI.Front
	l1 := config.MsgStateIntro
	l2 := ""
	if (credit != 0) || (len(self.inputBuf) > 0) {
		l1 = msgCredit + credit.FormatCtx(ctx)
		l2 = fmt.Sprintf(msgInputCode, string(self.inputBuf))
	}
	self.display.SetLines(l1, l2)
}
func (self *UI) onFrontSelectInput(ctx context.Context, e input.Event) State {
	switch {
	case e.IsDigit():
		self.inputBuf = append(self.inputBuf, byte(e.Key))
		return StateFrontSelect

	case input.IsReject(&e):
		// backspace semantic
		if len(self.inputBuf) > 0 {
			self.inputBuf = self.inputBuf[:len(self.inputBuf)-1]
			return StateFrontSelect
		}
		// TODO maybe ignore reject on empty buffer?
		return StateFrontEnd

	case input.IsAccept(&e):
		if len(self.inputBuf) == 0 {
			self.display.SetLines(self.g.Config.UI.Front.MsgError, msgMenuCodeEmpty)
			return StateInvalid
		}

		x, err := strconv.ParseUint(string(self.inputBuf), 10, 16)
		if err != nil {
			self.inputBuf = self.inputBuf[:0]
			self.display.SetLines(self.g.Config.UI.Front.MsgError, msgMenuCodeInvalid)
			return StateInvalid
		}
		code := uint16(x)

		mitem, ok := self.menu[code]
		if !ok {
			self.display.SetLines(self.g.Config.UI.Front.MsgError, msgMenuCodeInvalid)
			return StateInvalid
		}
		moneysys := money.GetGlobal(ctx)
		credit := moneysys.Credit(ctx)
		self.g.Log.Debugf("compare price=%v credit=%v", mitem.Price, credit)
		if mitem.Price > credit {
			self.display.SetLines(self.g.Config.UI.Front.MsgError, msgMenuInsufficientCredit)
			return StateInvalid
		}
		self.g.Log.Debugf("mitem=%s validate", mitem.String())
		if err := mitem.D.Validate(); err != nil {
			self.g.Log.Errorf("ui-front selected=%s Validate err=%v", mitem.String(), err)
			self.display.SetLines(self.g.Config.UI.Front.MsgError, msgMenuNotAvailable)
			return StateInvalid
		}

		self.FrontResult.Item = mitem
		return StateFrontAccept

	default:
		self.g.Log.Errorf("ui-front unhandled input=%v", e)
	}
	return StateFrontSelect
}

func (self *UI) onFrontTune(ctx context.Context) State {
	// XXX FIXME
	return self.onFrontSelect(ctx)
}

func (self *UI) onFrontTuneInput(e input.Event) State {
	switch e.Key {
	case input.EvendKeyCreamLess:
		if self.FrontResult.Cream > 0 {
			self.FrontResult.Cream--
			//lint:ignore SA9003 empty branch
		} else {
			// TODO notify "impossible input" (sound?)
		}
	case input.EvendKeyCreamMore:
		if self.FrontResult.Cream < MaxCream {
			self.FrontResult.Cream++
			//lint:ignore SA9003 empty branch
		} else {
			// TODO notify "impossible input" (sound?)
		}
	case input.EvendKeySugarLess:
		if self.FrontResult.Sugar > 0 {
			self.FrontResult.Sugar--
			//lint:ignore SA9003 empty branch
		} else {
			// TODO notify "impossible input" (sound?)
		}
	case input.EvendKeySugarMore:
		if self.FrontResult.Sugar < MaxSugar {
			self.FrontResult.Sugar++
			//lint:ignore SA9003 empty branch
		} else {
			// TODO notify "impossible input" (sound?)
		}
	default:
		return StateFrontSelect
	}
	var t1, t2 []byte
	next := StateFrontSelect
	switch e.Key {
	case input.EvendKeyCreamLess, input.EvendKeyCreamMore:
		t1 = self.display.Translate(fmt.Sprintf("%s  /%d", msgCream, self.FrontResult.Cream))
		t2 = formatScale(self.FrontResult.Cream, 0, MaxCream, ScaleAlpha)
		next = StateFrontTune
	case input.EvendKeySugarLess, input.EvendKeySugarMore:
		t1 = self.display.Translate(fmt.Sprintf("%s  /%d", msgSugar, self.FrontResult.Sugar))
		t2 = formatScale(self.FrontResult.Sugar, 0, MaxSugar, ScaleAlpha)
		next = StateFrontTune
	}
	t2 = append(append(append(make([]byte, 0, lcd.MaxWidth), '-', ' '), t2...), ' ', '+')
	self.display.SetLinesBytes(
		self.display.JustCenter(t1),
		self.display.JustCenter(t2),
	)
	return next
}

func (self *UI) onFrontAccept(ctx context.Context) State {
	moneysys := money.GetGlobal(ctx)
	uiConfig := &self.g.Config.UI
	selected := &self.FrontResult.Item
	teletx := tele.Telemetry_Transaction{
		Code:  int32(selected.Code),
		Price: uint32(selected.Price),
		// TODO options
		// TODO payment method
		// TODO bills, coins
	}
	self.g.Log.Debugf("ui-front selected=%s begin", selected.String())
	if err := moneysys.WithdrawPrepare(ctx, selected.Price); err != nil {
		self.g.Log.Errorf("ui-front CRITICAL error while return change")
	}
	itemCtx := money.SetCurrentPrice(ctx, selected.Price)
	if tuneCream := scaleTuneRate(self.FrontResult.Cream, MaxCream, DefaultCream); tuneCream != 1 {
		const name = "cream"
		var err error
		self.g.Log.Errorf("ui-front tuning stock=%s tune=%v", name, tuneCream)
		if itemCtx, err = self.g.Inventory.WithTuning(itemCtx, name, tuneCream); err != nil {
			self.g.Log.Errorf("ui-front tuning stock=%s err=%v", name, err)
		}
	}
	if tuneSugar := scaleTuneRate(self.FrontResult.Sugar, MaxSugar, DefaultSugar); tuneSugar != 1 {
		const name = "sugar"
		var err error
		self.g.Log.Errorf("ui-front tuning stock=%s tune=%v", name, tuneSugar)
		if itemCtx, err = self.g.Inventory.WithTuning(itemCtx, name, tuneSugar); err != nil {
			self.g.Log.Errorf("ui-front tuning stock=%s err=%v", name, err)
		}
	}
	self.display.SetLines(msgMaking1, msgMaking2)

	err := selected.D.Do(itemCtx)
	_ = self.g.Inventory.Persist.Store()
	self.g.Log.Debugf("ui-front selected=%s end err=%v", selected.String(), err)
	if err == nil { // success path
		self.g.Tele.Transaction(teletx)
		return StateFrontEnd
	}

	self.display.SetLines(uiConfig.Front.MsgError, uiConfig.Front.MsgMenuError)
	err = errors.Annotatef(err, "execute %s", selected.String())
	self.g.Error(err)

	if err := self.g.Engine.ExecList(ctx, "on_menu_error", self.g.Config.Engine.OnMenuError); err != nil {
		self.g.Error(err)
	} else {
		self.g.Log.Infof("on_menu_error success")
	}
	return StateBroken
}

func (self *UI) onFrontTimeout(ctx context.Context) State {
	self.g.Log.Debugf("ui state=%s result=%#v", self.State.String(), self.FrontResult)
	// moneysys := money.GetGlobal(ctx)
	// moneysys.save
	return StateFrontEnd
}

// tightly coupled to len(alphabet)=4
func formatScale(value, min, max uint8, alphabet []byte) []byte {
	var vicon [6]byte
	switch value {
	case min:
		vicon[0], vicon[1], vicon[2], vicon[3], vicon[4], vicon[5] = 0, 0, 0, 0, 0, 0
	case max:
		vicon[0], vicon[1], vicon[2], vicon[3], vicon[4], vicon[5] = 3, 3, 3, 3, 3, 3
	default:
		rng := uint16(max) - uint16(min)
		part := uint8((float32(value-min) / float32(rng)) * 24)
		// log.Printf("scale(%d,%d..%d) part=%d", value, min, max, part)
		for i := 0; i < len(vicon); i++ {
			if part >= 4 {
				vicon[i] = 3
				part -= 4
			} else {
				vicon[i] = part
				break
			}
		}
	}
	for i := 0; i < len(vicon); i++ {
		vicon[i] = alphabet[vicon[i]]
	}
	return vicon[:]
}

func scaleTuneRate(value, max, center uint8) float32 {
	switch {
	case value == center: // most common path
		return 1
	case value == 0:
		return 0
	}
	if value > max {
		value = max
	}
	if value > 0 && value < center {
		return 1 - (0.25 * float32(center-value))
	}
	if value > center && value <= max {
		return 1 + (0.25 * float32(value-center))
	}
	panic("code error")
}
