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

	exportStorage, err := exportFS.New(cfg.ExportPath, logger)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail init export fs",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	fileStorage, err := dataFS.New(cfg.FilePath, logger)
	if err != nil {
		logger.ErrorContext(
			ctx, "fail init data fs",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	// TODO: перейти со временем на юзкейсы
	c, err := api.New(
		time.Now(),
		logger,
		exportStorage,
		fileStorage,
		cfg.WebServerAddr,
		cfg.Debug,
		cfg.APIToken,
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
