package lib

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func ContextWithSignals(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	ready := make(chan struct{})

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		defer signal.Reset(os.Interrupt, syscall.SIGTERM)
		close(ready)

		select {
		case <-ctx.Done():
			return
		case <-sig:
			cancel()
		}
	}()

	<-ready
	return ctx
}
