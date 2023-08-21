package logger

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/exp/slog"
)

// SetupLogger sets up a new logger with the specified level and format based on whether the application is running locally or not
func SetupLogger(isLocal bool, level string) (*slog.Logger, error) {
	l, err := newLevelFromString(level)
	if err != nil {
		return nil, err
	}
	opts := &slog.HandlerOptions{Level: l}

	// Create a new handler based on whether the application is running locally or not
	var handler slog.Handler
	if isLocal {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	return slog.New(handler), nil
}

type contextKeyType struct{}

var contextKey = contextKeyType{}

// WithLogger adds a logger to the context
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, contextKey, logger)
}

// FromContext retrieves the logger from the context
func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(contextKey).(*slog.Logger); ok {
		return logger
	}
	return slog.Default()
}

// newLevelFromString converts a level string to a slog.Level
func newLevelFromString(level string) (slog.Level, error) {
	switch level {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	}

	return slog.Level(0), fmt.Errorf("unknown level: %q", level)
}
