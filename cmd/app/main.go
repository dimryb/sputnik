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
	"github.com/dimryb/sputnik/internal/service"
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
	application := app.NewApp(ctx, logg)
	sputnikService := service.NewSputnikService(application, logg, service.SputnikConfig{
		Host:              cfg.HTTP.Host,
		Port:              cfg.HTTP.Port,
		ReadTimeout:       cfg.HTTP.ReadTimeout,
		WriteTimeout:      cfg.HTTP.WriteTimeout,
		IdleTimeout:       cfg.HTTP.IdleTimeout,
		ReadHeaderTimeout: cfg.HTTP.ReadHeaderTimeout,
	})

	if err = sputnikService.Run(ctx); err != nil {
		logg.Errorf("Sputnik service stopped with error: %v", err)
		cancel()
	} else {
		logg.Infof("Sputnik service stopped gracefully")
	}

	<-ctx.Done()

	fmt.Println("App stopped")
}
