package service

import (
	"context"

	i "github.com/dimryb/sputnik/internal/interface"
)

type SputnikConfig struct{}

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
	_ = ctx

	return nil
}
