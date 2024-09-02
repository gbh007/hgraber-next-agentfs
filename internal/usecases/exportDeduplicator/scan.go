package exportDeduplicator

import (
	"context"
	"fmt"
	"hgnextfs/internal/entities"
	"log/slog"
	"time"
)

const minEntryPercentage = 0.9999 // Считаем допустимой погрешностью 1 страницу на 10 000.

func (uc *UseCase) ScanZips(ctx context.Context) error {
	relativePaths, err := uc.exportFS.AllZips(ctx)
	if err != nil {
		return fmt.Errorf("export fs: scan all zip: %w", err)
	}

	for i, relativePath := range relativePaths {
		uc.logger.DebugContext(
			ctx, "start match archive",
			slog.Int("current", i+1),
			slog.Int("total", len(relativePaths)),
			slog.String("path", relativePath),
		)

		body, err := uc.exportFS.Get(ctx, relativePath)
		if err != nil {
			return fmt.Errorf("export fs: get zip body (%s): %w", relativePath, err)
		}

		matches, err := uc.masterAPI.DeduplicateArchive(ctx, body)
		if err != nil {
			return fmt.Errorf("master api match (%s): %w", relativePath, err)
		}

		var (
			matched            bool
			maxEntryPercentage float64
		)

		for _, match := range matches {
			if match.EntryPercentage > minEntryPercentage &&
				match.ReverseEntryPercentage > minEntryPercentage {
				err = uc.storage.CreateExport(ctx, entities.ExportInfo{
					BookID:     match.TargetBookID,
					BookURL:    match.OriginBookURL,
					FSPath:     relativePath,
					ExportedAt: time.Now().UTC(),
				})
				if err != nil {
					return fmt.Errorf("storage create export info (%s): %w", relativePath, err)
				}

				matched = true
			}

			if match.EntryPercentage > maxEntryPercentage {
				maxEntryPercentage = match.EntryPercentage
			}
		}

		if !matched {
			err = uc.storage.CreateMissing(ctx, relativePath, maxEntryPercentage)
			if err != nil {
				return fmt.Errorf("storage create missing info (%s): %w", relativePath, err)
			}
		}
	}

	return nil
}
