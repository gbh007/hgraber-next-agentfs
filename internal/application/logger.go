package application

import (
	"context"
	"hgnextfs/internal/config"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel/trace"
)

func initLogger(cfg config.Config) *slog.Logger {
	slogOpt := &slog.HandlerOptions{
		AddSource: cfg.Application.Debug,
		Level:     slog.LevelInfo,
	}

	if cfg.Application.Debug {
		slogOpt.Level = slog.LevelDebug
	}

	return slog.New(
		logHandler{
			Handler: slog.NewJSONHandler(
				os.Stderr,
				slogOpt,
			),
		},
	)
}

// TODO: в случае использования групп реализовать более безопасно.
type logHandler struct {
	slog.Handler
}

func (lh logHandler) Handle(ctx context.Context, r slog.Record) error {
	snapContext := trace.SpanContextFromContext(ctx)
	if snapContext.HasTraceID() {
		r.AddAttrs(slog.String("trace_id", snapContext.TraceID().String()))
	}

	return lh.Handler.Handle(ctx, r)
}

func (lh logHandler) WithGroup(name string) slog.Handler {
	return logHandler{
		Handler: lh.Handler.WithGroup(name),
	}
}

func (lh logHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return logHandler{
		Handler: lh.Handler.WithAttrs(attrs),
	}
}
