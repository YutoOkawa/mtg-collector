package main

import (
	"context"
	"fmt"
	"mtg-collector/pkg/server"
	"os/signal"
	"syscall"
)

func main() {
	server := server.NewServer(8080)
	ctx, ctxCancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer ctxCancel()

	srvErrCh := make(chan error, 1)
	go server.Start(srvErrCh)

	for {
		select {
		case err := <-srvErrCh:
			fmt.Println(err)
		case <-ctx.Done():
			fmt.Println("Graceful shutdown")
		}
	}
}
