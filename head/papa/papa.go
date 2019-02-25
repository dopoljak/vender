package papa

import (
	"context"
	"net"
	"time"

	"github.com/temoto/vender/head/state"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//go:generate protoc -I=../../protobuf --go_out=plugins=grpc:./ ../../protobuf/papa.proto

func (self *PapaSystem) netLoop(ctx context.Context) {
	// TODO alive
	for self.alive.IsRunning() {
		client, err := self.dial(ctx)
		if err == nil {
			err = self.network(ctx, client)
		}
		if err != nil {
			self.Log.Errorf(err.Error())
			time.Sleep(1 * time.Second)
			if neterr, ok := err.(net.Error); ok {
				if neterr.Temporary() {
					continue
				} else {
					self.Log.Errorf("network error is permanent")
					return
				}
			}
		}
	}
}

func (self *PapaSystem) network(ctx context.Context, client PapaClient) error {
	pull, err := client.Pull(ctx)
	if err != nil {
		return err
	}
	for {
		task, err := pull.Recv()
		if err != nil {
			self.Log.Errorf(err.Error())
			// TODO handle error
			time.Sleep(1 * time.Second)
			continue
		}

		// TODO execute later from on-disk queue
		switch task.GetName() {
		case "restart-head":
			go time.AfterFunc(5*time.Second, func() {
				l := ctx.Value("lifecycle").(*state.Lifecycle)
				l.Restart(ctx)
			})
		}

		// TODO write task to disk, then execute
		err = pull.Send(&PapaTask{Id: task.Id})
		// TODO handle error
		_ = err
	}
}

func (self *PapaSystem) dial(ctx context.Context) (PapaClient, error) {
	config := state.GetConfig(ctx)

	optSecurity := grpc.WithInsecure()
	if config.Papa.CertFile != "" {
		creds, err := credentials.NewClientTLSFromFile(config.Papa.CertFile, "")
		if err != nil {
			self.Log.Errorf(err.Error())
			return nil, err
		}
		optSecurity = grpc.WithTransportCredentials(creds)
	}

	conn, err := grpc.Dial(config.Papa.Address, optSecurity)
	if err != nil {
		self.Log.Errorf(err.Error())
		return nil, err
	}

	client := NewPapaClient(conn)
	return client, nil
}
