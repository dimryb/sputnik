package app

import (
	"context"
	"fmt"

	i "github.com/dimryb/sputnik/internal/interface"
)

type App struct {
	ctx    context.Context
	Logger i.Logger
}

func NewApp(ctx context.Context, logger i.Logger) *App {
	return &App{
		ctx:    ctx,
		Logger: logger,
	}
}

func (a *App) Run() {
	fmt.Println("App started")
	select { //nolint: gosimple
	case <-a.ctx.Done():
	}
}
