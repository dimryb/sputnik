package service

import (
	"context"
	"fmt"
	"time"

	i "github.com/dimryb/sputnik/internal/interface"
	internalhttp "github.com/dimryb/sputnik/internal/server/http"
)

type SputnikConfig struct {
	Host              string
	Port              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	ReadHeaderTimeout time.Duration
}

type Sputnik struct {
	app  i.Application
	logg i.Logger
	cfg  SputnikConfig
}

func NewSputnikService(app i.Application, logg i.Logger, cfg SputnikConfig) *Sputnik {
	return &Sputnik{
		app:  app,
		logg: logg,
		cfg:  cfg,
	}
}

func (s *Sputnik) Run(ctx context.Context) error {
	handlers := internalhttp.NewHandlers(s.app, s.logg)
	server := internalhttp.NewServer(s.app, s.logg, internalhttp.ServerConfig{
		Host:              s.cfg.Host,
		Port:              s.cfg.Port,
		ReadTimeout:       s.cfg.ReadTimeout,
		WriteTimeout:      s.cfg.WriteTimeout,
		IdleTimeout:       s.cfg.IdleTimeout,
		ReadHeaderTimeout: s.cfg.ReadHeaderTimeout,
	}, handlers)

	go func() {
		<-ctx.Done()
		s.logg.Infof("Stopping HTTP server...")
		if err := server.Stop(context.Background()); err != nil {
			s.logg.Errorf("Failed to stop http server: %s", err.Error())
		}
	}()

	s.logg.Infof("sputnik is running...")

	if err := server.Start(ctx); err != nil {
		return fmt.Errorf("failed to start http server: %s", err.Error())
	}
	return nil
}
