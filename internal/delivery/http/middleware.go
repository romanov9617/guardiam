package http

import (
	"log/slog"
	"net/http"
	"time"
)

// responseWriter wrapper for capturing status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// LoggingMiddleware returns middleware, which logs each request
// and response via slog.Logger.
func LoggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			next.ServeHTTP(rw, r)

			duration := time.Since(start)
			logger.Info("HTTP request",
				slog.String("method", r.Method),
				slog.String("url", r.URL.RequestURI()),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
				slog.Time("started_at", start),
				slog.Duration("duration", duration),
				slog.Int("status", rw.statusCode),
			)
		})
	}
}
