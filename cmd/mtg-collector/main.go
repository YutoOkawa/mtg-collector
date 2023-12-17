package main

import (
	"context"
	"mtg-collector/pkg/logger"
	"mtg-collector/pkg/server"
	"os/signal"
	"syscall"
)

func main() {
	server := server.NewServer(8080, 1)
	ctx, ctxCancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer ctxCancel()

	log := logger.GetLogger()

	srvErrCh := make(chan error, 1)
	go server.Start(srvErrCh)

	for {
		select {
		case err := <-srvErrCh:
			log.Error("failed to initialize server",
				logger.String("err", err.Error()))
			return
		case <-ctx.Done():
			if err := server.Shutdown(); err != nil {
				log.Error("failed to shutdown server",
					logger.String("err", err.Error()))
			}
			log.Info("shutdown")
			return
		}
	}
}
