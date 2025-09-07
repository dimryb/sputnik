package app

import (
	"context"
	"fmt"
)

type App struct {
	ctx context.Context
}

func NewApp(ctx context.Context) *App {
	return &App{ctx: ctx}
}

func (a *App) Run() {
	fmt.Println("App started")
	select { //nolint: gosimple
	case <-a.ctx.Done():
	}
}
