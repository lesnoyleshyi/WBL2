package main

import (
	"WBL2/Tasks/T11/internal/app"
	"context"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	defer stop()

	go app.Start(ctx)

	<-ctx.Done()
	app.Stop(ctx)
}
