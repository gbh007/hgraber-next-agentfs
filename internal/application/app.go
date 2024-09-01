package application

import (
	"context"
	"hgnextfs/internal/adapter/dataFS"
	"hgnextfs/internal/adapter/exportFS"
	"hgnextfs/internal/controller/api"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Serve() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	cfg, err := parseConfig()
	if err != nil {
		// Поскольку на этот момент нет ни логгера ни вообще ничего то выкидываем панику.
		panic(err)
	}

	logger := initLogger(cfg)

	if cfg.Application.TraceEndpoint != "" {
		err := initTrace(ctx, cfg.Application.TraceEndpoint)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init otel",
				slog.Any("error", err),
			)

			os.Exit(1)
		}
	}

	var (
		exportStorage *exportFS.Storage
		fileStorage   *dataFS.Storage
	)

	if cfg.FSBase.ExportPath != "" {
		exportStorage, err = exportFS.New(cfg.FSBase.ExportPath, logger)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init export fs",
				slog.Any("error", err),
			)

			os.Exit(1)
		}

		logger.DebugContext(
			ctx, "use local export storage",
			slog.String("path", cfg.FSBase.ExportPath),
		)
	}

	if cfg.FSBase.FilePath != "" {
		fileStorage, err = dataFS.New(cfg.FSBase.FilePath, logger)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init data fs",
				slog.Any("error", err),
			)

			os.Exit(1)
		}

		logger.DebugContext(
			ctx, "use local file storage",
			slog.String("path", cfg.FSBase.FilePath),
		)
	}

	// TODO: перейти со временем на юзкейсы
	c, err := api.New(
		time.Now(),
		logger,
		exportStorage,
		fileStorage,
		cfg.API.Addr,
		cfg.Application.Debug,
		cfg.API.Token,
	)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail init api controller",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	logger.InfoContext(ctx, "application start")
	defer logger.InfoContext(ctx, "application stop")

	err = c.Serve(ctx)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail serve api",
			slog.Any("error", err),
		)

		os.Exit(1)
	}
}
