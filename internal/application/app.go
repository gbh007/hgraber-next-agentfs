package application

import (
	"context"
	"hgnextfs/internal/adapter/dataFS"
	"hgnextfs/internal/adapter/exportFS"
	"hgnextfs/internal/adapter/masterAPI"
	"hgnextfs/internal/adapter/storage"
	"hgnextfs/internal/controller/api"
	"hgnextfs/internal/usecases/exportAPI"
	"hgnextfs/internal/usecases/exportDeduplicator"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.opentelemetry.io/otel"
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

	cfg, needScan, err := parseConfig()
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

	tracer := otel.GetTracerProvider().Tracer("hgraber-next")

	var (
		exportStorage api.ExportUseCase
		fileStorage   api.FileUseCase

		exportStorageRaw *exportFS.Storage
		dbRaw            *storage.Storage
		mAPI             *masterAPI.Client
	)

	if cfg.FSBase.ExportPath != "" {
		exportStorageRaw, err = exportFS.New(cfg.FSBase.ExportPath, logger, cfg.FSBase.ExportLimitOnFolder)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init export fs",
				slog.Any("error", err),
			)

			os.Exit(1)
		}

		exportStorage = exportStorageRaw

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

	if cfg.Sqlite.FilePath != "" {
		dbRaw, err = storage.New(ctx, logger, cfg.Sqlite.FilePath)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init db",
				slog.Any("error", err),
			)

			os.Exit(1)
		}
	}

	if cfg.ZipScanner.MasterAddr != "" {
		mAPI, err = masterAPI.New(cfg.ZipScanner.MasterAddr, cfg.ZipScanner.MasterToken)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail init master api",
				slog.Any("error", err),
			)

			os.Exit(1)
		}
	}

	if needScan {
		if dbRaw == nil || exportStorageRaw == nil || mAPI == nil {
			logger.ErrorContext(ctx, "invalid scan dependencies")

			os.Exit(1)
		}

		err = exportDeduplicator.New(logger, exportStorageRaw, dbRaw, mAPI).ScanZips(ctx)
		if err != nil {
			logger.ErrorContext(
				ctx, "fail scan zips",
				slog.Any("error", err),
			)

			os.Exit(1)
		}

		return
	}

	if cfg.FSBase.EnableDeduplication && dbRaw != nil && exportStorageRaw != nil {
		exportStorage = exportAPI.New(logger, dbRaw, exportStorageRaw)

		logger.DebugContext(ctx, "use export deduplication")
	}

	// TODO: перейти со временем на юзкейсы
	c, err := api.New(
		time.Now(),
		logger,
		tracer,
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
