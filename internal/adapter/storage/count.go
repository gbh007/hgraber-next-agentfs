package storage

import (
	"context"
	"database/sql"
	"net/url"

	"github.com/google/uuid"
)

func (s *Storage) ExportedCountByID(ctx context.Context, bookID uuid.UUID) (int, error) {
	var c sql.NullInt64

	err := s.db.GetContext(ctx, &c, `SELECT COUNT(*) FROM export_infos WHERE book_id = ?;`, bookID)
	if err != nil {
		return 0, err
	}

	return int(c.Int64), nil
}

func (s *Storage) ExportedCountByURL(ctx context.Context, u url.URL) (int, error) {
	var c sql.NullInt64

	err := s.db.GetContext(ctx, &c, `SELECT COUNT(*) FROM export_infos WHERE book_url = ?;`, u.String())
	if err != nil {
		return 0, err
	}

	return int(c.Int64), nil
}
