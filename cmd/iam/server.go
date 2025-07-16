package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cmd/iam/main.go/internal/config"
	api "cmd/iam/main.go/internal/delivery/http"
	"cmd/iam/main.go/internal/logging"
)

type Server struct {
	db     *sql.DB
	server *http.Server
	cfg    *config.Config
	logger *slog.Logger
}

func NewServer(cfg *config.Config) (*Server, error) {

	logger := logging.SetUpLogger(cfg.Logger.Environment)
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%v/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("db open: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("db doesn't ping")
	}
	logger.Info("Database connection successfull")

	server, err := api.ConfigureHTTPServer(cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("config http server: %w", err)
	}
	return &Server{
		db:     db,
		server: server,
		cfg:    cfg,
		logger: logger,
	}, err
}

func (s *Server) Run() <-chan error {
	errChan := make(chan error, 1)
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT,
	)

	go func() {
		<-ctx.Done()
		s.logger.Info("Shutting down...")
		ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(s.cfg.API.ShutdownIn)*time.Second)
		if err := s.server.Shutdown(ctxTimeout); err != nil {
			errChan <- err
		}

		s.logger.Info("Successfully shutting down")
		defer func() {
			stop()
			cancel()
			close(errChan)
		}()
	}()

	go func() {
		s.logger.Info("Starting server...")
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()

	return errChan
}
