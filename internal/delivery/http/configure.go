// Package http provides a configurable HTTP server with middleware support.
//
// Package http allows configuring an HTTP server and attaching middleware
// in a reusable and composable way. Middleware can wrap individual handlers
// or the entire server, enabling cross-cutting concerns such as logging,
// authentication, request timeout, and metrics.
package http

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	"cmd/iam/main.go/internal/config"
)

func ConfigureHTTPServer(cfg *config.Config, logger *slog.Logger) (*http.Server, error) {
	addr := fmt.Sprintf("%s:%v", cfg.API.Host, cfg.API.Port)

	loggingMiddleware := LoggingMiddleware(logger)
	r := chi.NewRouter()
	r.Use(loggingMiddleware)

	server := &http.Server{
		Handler: r,
		Addr:    addr,
	}
	return server, nil
}
