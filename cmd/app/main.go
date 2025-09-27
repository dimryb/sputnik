package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/dimryb/sputnik/internal/app"
	"github.com/dimryb/sputnik/internal/config"
	"github.com/dimryb/sputnik/internal/logger"
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

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	logg := logger.New(cfg.Log.Level)

	app.NewApp(ctx, logg).Run()

	<-ctx.Done()

	fmt.Println("App stopped")
}
