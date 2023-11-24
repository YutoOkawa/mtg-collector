package main

import (
	"context"
	"fmt"
	"mtg-collector/pkg/server"
	"os/signal"
	"syscall"
)

func main() {
	server := server.NewServer(8080, 1)
	ctx, ctxCancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer ctxCancel()

	srvErrCh := make(chan error, 1)
	go server.Start(srvErrCh)

	for {
		select {
		case err := <-srvErrCh:
			fmt.Println(err)
		case <-ctx.Done():
			if err := server.Shutdown(); err != nil {
				// TODO: log 出力
			}
			return
		}
	}
}
