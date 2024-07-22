package dataFS

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"

	"hgnextfs/internal/entities"
)

func (s *Storage) Delete(ctx context.Context, fileID uuid.UUID) error {
	filepath := s.filepath(fileID)

	err := os.Remove(filepath)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("local fs: %w", entities.FileNotFoundError)
	}

	if err != nil {
		return fmt.Errorf("local fs: os remove: %w", err)
	}

	return nil
}
