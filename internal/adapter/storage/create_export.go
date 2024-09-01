package storage

import (
	"context"
	"hgnextfs/internal/entities"
)

func (s *Storage) CreateExport(ctx context.Context, info entities.ExportInfo) error {
	_, err := s.db.ExecContext(
		ctx,
		`INSERT INTO export_infos (book_id, book_url, relative_path, exported_at) VALUES (?,?,?,?) ON CONFLICT DO NOTHING;`,
		info.BookID,
		URLToDB(info.BookURL),
		info.FSPath,
		info.ExportedAt.Unix(),
	)
	if err != nil {
		return err
	}

	return nil
}
