package internalhttp

import (
	"net/http"

	i "github.com/dimryb/sputnik/internal/interface"
)

type Handlers struct {
	app    i.Application
	logger i.Logger
}

func NewHandlers(app i.Application, logger i.Logger) *Handlers {
	return &Handlers{app, logger}
}

func (h *Handlers) helloHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}
