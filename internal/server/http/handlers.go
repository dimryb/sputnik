package internalhttp

import (
	"net/http"

	i "github.com/dimryb/sputnik/internal/interface"
)

type CalendarHandlers struct {
	app    i.Application
	logger i.Logger
}

func NewCalendarHandlers(app i.Application, logger i.Logger) *CalendarHandlers {
	return &CalendarHandlers{app, logger}
}

func (h *CalendarHandlers) helloHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}
