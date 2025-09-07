package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/dimryb/sputnik/internal/app"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "configs/config.yaml", "Path to configuration file")
}

func main() {
	fmt.Println("Hello World, Sputnik")

	flag.Parse()

	printVersion()
	if flag.Arg(0) == "version" {
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	app.NewApp(ctx).Run()

	<-ctx.Done()

	fmt.Println("App stopped")
}
