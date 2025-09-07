package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/dimryb/sputnik/internal/app"
)

func main() {
	fmt.Println("Hello World, Sputnik")

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	app.NewApp(ctx).Run()

	<-ctx.Done()

	fmt.Println("App stopped")
}
