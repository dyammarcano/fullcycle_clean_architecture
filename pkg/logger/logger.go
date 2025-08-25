package logger

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/dyammarcano/fullcycle_clean_architecture/pkg/config"
)

func init() {
	opts := &slog.HandlerOptions{
		Level: setLevel(),
	}

	logger := slog.New(setFormat(os.Stdout, opts))
	slog.SetDefault(logger)
}

func LogContext(ctx context.Context, level slog.Level, msg string, args ...any) {
	slog.Log(ctx, level, msg, args...)
}

func Log(level slog.Level, msg string, args ...any) {
	LogContext(context.Background(), level, msg, args...)
}

func setLevel() slog.Level {
	switch config.G.Logger.LogLevel {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func setFormat(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	switch config.G.Logger.LogFormat {
	case "json":
		return slog.NewJSONHandler(w, opts)
	default:
		return slog.NewTextHandler(w, opts)
	}
}
