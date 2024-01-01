package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type Server struct {
	*http.Server
	logger *zap.Logger
	router *http.ServeMux
}

// Returns a new Server struct with a logger
func NewServer(logger *zap.Logger) *Server {
	logger.Info("Reading config")
	config, err := NewConfig()
	if err != nil {
		logger.Fatal(err.Error())
	}

	router := http.NewServeMux()

	srv := &http.Server{
		Addr:        config.Address,
		Handler:     router,
		ReadTimeout: time.Duration(config.ReadTimeout) * time.Second,
		IdleTimeout: time.Duration(config.IdleTimeout) * time.Second,
	}

	s := &Server{
		Server: srv,
		logger: logger,
		router: router,
	}

	return s
}

// Handles GET requests
func (s *Server) GET(pattern string, handler func(*Context)) {
	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := Context{Writer: w, Request: r}
		if r.Method != http.MethodGet {
			ctx.Error(http.StatusBadRequest, fmt.Errorf("request method must be GET"))
			return
		}
		ctx.ParseParams()
		handler(&ctx)
	}
	s.router.HandleFunc(pattern, Logger(h))
}

// Handles POST requests
func (s *Server) POST(pattern string, handler func(*Context)) {
	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := Context{Writer: w, Request: r}
		if r.Method != http.MethodPost {
			ctx.Error(http.StatusBadRequest, fmt.Errorf("request method must be POST"))
			return
		}
		ctx.ParseParams()
		handler(&ctx)
	}
	s.router.HandleFunc(pattern, Logger(h))
}

// Runs the server
func (s *Server) Run() {
	s.logger.Info("Starting HTTP server")
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("Can't start server", zap.Error(err), zap.String("Server address", s.Addr))
		}
	}()
	s.logger.Info(fmt.Sprintf("Listening on %v", s.Addr))

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	s.logger.Info("Interrupt, shutting down the server")
	timeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()
	if err := s.Shutdown(timeout); err != nil {
		s.logger.Error("Can't shut down the server", zap.Error(err))
	}

}
