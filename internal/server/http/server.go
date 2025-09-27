package internalhttp

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	i "github.com/dimryb/sputnik/internal/interface"
)

type Server struct {
	app    i.Application
	logger i.Logger
	server *http.Server
	cfg    ServerConfig
}

type ServerConfig struct {
	Host              string
	Port              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	ReadHeaderTimeout time.Duration
}

func NewServer(app i.Application, logger i.Logger, cfg ServerConfig, handlers *Handlers) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.helloHandler)
	return &Server{
		logger: logger,
		app:    app,
		server: &http.Server{
			Handler:           loggingMiddleware(handlers.logger)(mux),
			ReadTimeout:       cfg.ReadTimeout,
			WriteTimeout:      cfg.WriteTimeout,
			IdleTimeout:       cfg.IdleTimeout,
			ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		},
		cfg: cfg,
	}
}

func (s *Server) Start(_ context.Context) error {
	addr := net.JoinHostPort(s.cfg.Host, s.cfg.Port)
	s.server.Addr = addr

	s.logger.Infof(fmt.Sprintf("Starting HTTP server on %s", addr))
	err := s.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Errorf("Failed to start HTTP server: " + err.Error())
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Infof("Stopping HTTP server")
	return s.server.Shutdown(ctx)
}

func (s *Server) Handler() http.Handler {
	return s.server.Handler
}
