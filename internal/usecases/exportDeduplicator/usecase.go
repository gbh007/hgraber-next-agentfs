package exportDeduplicator

import (
	"context"
	"hgnextfs/internal/entities"
	"io"
	"log/slog"
)

type exportFS interface {
	AllZips(ctx context.Context) ([]string, error)
	Get(ctx context.Context, relativePath string) (io.Reader, error)
}

type storage interface {
	CreateExport(ctx context.Context, info entities.ExportInfo) error
	CreateMissing(ctx context.Context, path string, maxEntryPercentage float64) error
}

type masterAPI interface {
	DeduplicateArchive(ctx context.Context, body io.Reader) ([]entities.DeduplicateArchiveResult, error)
}

type UseCase struct {
	logger *slog.Logger

	exportFS  exportFS
	storage   storage
	masterAPI masterAPI
}

func New(
	logger *slog.Logger,
	fsScanner exportFS,
	storage storage,
	masterAPI masterAPI,
) *UseCase {
	return &UseCase{
		logger:    logger,
		exportFS:  fsScanner,
		storage:   storage,
		masterAPI: masterAPI,
	}
}
