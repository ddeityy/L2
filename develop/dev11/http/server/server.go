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

func (s *Server) AddRoute(pattern string, handler http.HandlerFunc) {
	s.router.HandleFunc(pattern, Logger(handler))
}

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
	s.logger.Info("Received interrupt signal, shutting down in 5 seconds")

	timeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()
	if err := s.Shutdown(timeout); err != nil {
		s.logger.Error("Can't close server", zap.Error(err))
	}
}
