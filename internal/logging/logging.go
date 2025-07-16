// Package logging provides a simple way to set up a logger
// with sensible defaults depending on the application environment.
//
// The logger output format and log level are selected based on the environment value.
//
// Example usage:
//
//	logger := logging.SetUpLogger("dev")
//	logger.Info("Application started")
package logging

import (
	"log/slog"
	"os"
)

const (
	envDev   = "dev"
	envProd  = "prod"
	envLocal = "local"
)

// SetUpLogger initializes and returns a *slog.Logger configured for the specified environment.
//
// For "dev", the logger outputs JSON with Debug level.
// For "prod", the logger outputs JSON with Info level.
// For "local" and any other value, the logger outputs plain text with Debug level.
func SetUpLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	case envLocal:
		fallthrough
	default:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}
	return log
}
