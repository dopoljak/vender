// Temporary executable for developing UI
package main

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/temoto/vender/engine"
	"github.com/temoto/vender/head/money"
	"github.com/temoto/vender/head/tele"
	"github.com/temoto/vender/head/ui"
	"github.com/temoto/vender/log2"
	"github.com/temoto/vender/state"
)

var log = log2.NewStderr(log2.LDebug)

func main() {
	cmdline := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagConfig := cmdline.String("config", "vender.hcl", "")
	cmdline.Parse(os.Args[1:])

	log.SetFlags(log2.LInteractiveFlags)

	ctx := context.Background()
	ctx = context.WithValue(ctx, log2.ContextKey, log)
	ctx = context.WithValue(ctx, engine.ContextKey, engine.NewEngine(ctx))
	config := state.MustReadConfigFile(ctx, *flagConfig)
	log.Debugf("config=%+v", config)
	ctx = state.ContextWithConfig(ctx, config)

	log.Debugf("Init display")
	d := config.Global().Hardware.HD44780.Display
	d.SetLine1("loaded")

	moneysys := new(money.MoneySystem)
	moneysys.Start(ctx)
	telesys := config.Global().Tele

	menuMap := make(ui.Menu)
	menuMap.Add(1, "chai", config.ScaleU(3),
		engine.Func0{F: func() error {
			d.SetLines("спасибо", "готовим...")
			time.Sleep(7 * time.Second)
			d.SetLines("успех", "спасибо")
			time.Sleep(3 * time.Second)
			return nil
		}})
	menuMap.Add(2, "coffee", config.ScaleU(5),
		engine.Func0{F: func() error {
			d.SetLines("спасибо", "готовим...")
			time.Sleep(7 * time.Second)
			d.SetLines("успех", "спасибо")
			time.Sleep(3 * time.Second)
			return nil
		}})

	for {
		menu := ui.NewUIMenu(ctx, menuMap)
		menu.SetCredit(moneysys.Credit(ctx))
		moneysys.AcceptCredit(ctx, menuMap.MaxPrice())

		stopCh := menu.StopChan()
		go func() {
			for {
				select {
				case <-stopCh:
					return
				case em := <-moneysys.Events():
					log.Debugf("money event: %s", em.String())
					switch em.Name() {
					case money.EventCredit:
						menu.SetCredit(moneysys.Credit(ctx))
						moneysys.AcceptCredit(ctx, menuMap.MaxPrice())
					case money.EventAbort:
						err := moneysys.Abort(ctx)
						log.Infof("user requested abort err=%v", err)
						menu.SetCredit(moneysys.Credit(ctx))
						moneysys.AcceptCredit(ctx, menuMap.MaxPrice())
					default:
						panic("head: unknown money event: " + em.String())
					}
				case cmd := <-telesys.CommandChan():
					switch cmd.Task {
					case tele.Command_Abort:
						err := moneysys.Abort(ctx)
						telesys.CommandReplyErr(cmd, err)
						log.Infof("admin requested abort err=%v", err)
						menu.SetCredit(moneysys.Credit(ctx))
						moneysys.AcceptCredit(ctx, menuMap.MaxPrice())
					}
				}
			}
		}()

		result := menu.Run()
		log.Debugf("result=%#v", result)
		if !result.Confirm {
			continue
		}

		err := moneysys.WithdrawPrepare(ctx, result.Item.Price)
		if err == money.ErrNeedMoreMoney {
			log.Errorf("menuitem=%v price=%s err=%v", result.Item, result.Item.Price.FormatCtx(ctx), err)
		} else if err == nil {
			if err = result.Item.D.Do(ctx); err != nil {
				log.Errorf("menuitem=%v execute err=%v", result.Item, err)
				moneysys.Abort(ctx)
			} else {
				moneysys.WithdrawCommit(ctx, result.Item.Price)
			}
		} else {
			log.Errorf("menuitem=%v price=%s err=%v", result.Item, result.Item.Price.FormatCtx(ctx), err)
		}
	}
}