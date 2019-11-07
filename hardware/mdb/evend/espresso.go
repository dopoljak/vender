package evend

import (
	"context"
	"fmt"
	"time"

	"github.com/juju/errors"
	"github.com/temoto/vender/engine"
	"github.com/temoto/vender/helpers"
	"github.com/temoto/vender/state"
)

const DefaultEspressoTimeout = 30 * time.Second

type DeviceEspresso struct {
	Generic

	timeout time.Duration
}

func (self *DeviceEspresso) Init(ctx context.Context) error {
	g := state.GetGlobal(ctx)
	espressoConfig := &g.Config.Hardware.Evend.Espresso
	self.timeout = helpers.IntSecondDefault(espressoConfig.TimeoutSec, DefaultEspressoTimeout)
	err := self.Generic.Init(ctx, 0xe8, "espresso", proto2)
	if err != nil {
		return errors.Annotate(err, "evend.espresso.Init")
	}

	g.Engine.Register("mdb.evend.espresso_grind", self.Generic.WithRestart(self.NewGrind()))
	g.Engine.Register("mdb.evend.espresso_press", self.NewPress())
	g.Engine.Register("mdb.evend.espresso_dispose", self.Generic.WithRestart(self.NewRelease()))
	g.Engine.Register("mdb.evend.espresso_heat_on", self.NewHeat(true))
	g.Engine.Register("mdb.evend.espresso_heat_off", self.NewHeat(false))

	return nil
}

func (self *DeviceEspresso) NewGrind() engine.Doer {
	const tag = "mdb.evend.espresso.grind"
	return engine.NewSeq(tag).
		Append(self.Generic.NewWaitReady(tag)).
		Append(self.Generic.NewAction(tag, 0x01)).
		// TODO expect delay like in cup dispense, ignore immediate error, retry
		Append(self.Generic.NewWaitDone(tag, self.timeout))
}

func (self *DeviceEspresso) NewPress() engine.Doer {
	const tag = "mdb.evend.espresso.press"
	return engine.NewSeq(tag).
		Append(self.Generic.NewWaitReady(tag)).
		Append(self.Generic.NewAction(tag, 0x02)).
		Append(self.Generic.NewWaitDone(tag, self.timeout))
}

func (self *DeviceEspresso) NewRelease() engine.Doer {
	const tag = "mdb.evend.espresso.release"
	return engine.NewSeq(tag).
		Append(self.Generic.NewWaitReady(tag)).
		Append(self.Generic.NewAction(tag, 0x03)).
		Append(self.Generic.NewWaitDone(tag, self.timeout))
}

func (self *DeviceEspresso) NewHeat(on bool) engine.Doer {
	tag := fmt.Sprintf("mdb.evend.espresso.heat:%t", on)
	arg := byte(0x05)
	if !on {
		arg = 0x06
	}
	return engine.NewSeq(tag).
		Append(self.Generic.NewWaitReady(tag)).
		Append(self.Generic.NewAction(tag, arg))
}
