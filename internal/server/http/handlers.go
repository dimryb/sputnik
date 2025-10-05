package internalhttp

import (
	"encoding/json"
	"net/http"
	"time"

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

func (h *Handlers) healthHandler(w http.ResponseWriter, _ *http.Request) {
	type HealthResponse struct {
		Status    string    `json:"status"`
		Timestamp time.Time `json:"timestamp"`
	}

	resp := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Errorf("failed to encode health response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
