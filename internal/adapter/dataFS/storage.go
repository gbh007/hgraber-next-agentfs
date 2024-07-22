package dataFS

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/google/uuid"
)

type Storage struct {
	fsPath string

	logger *slog.Logger
}

func New(path string, logger *slog.Logger) (*Storage, error) {
	err := createDir(path)
	if err != nil {
		return nil, err
	}

	return &Storage{
		fsPath: path,
		logger: logger,
	}, nil
}

func (s *Storage) filepath(fileID uuid.UUID) string {
	return path.Join(s.fsPath, fileID.String())
}

func createDir(dirPath string) error {
	info, err := os.Stat(dirPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if info != nil && !info.IsDir() {
		return fmt.Errorf("dir path is not dir")
	}

	err = os.MkdirAll(dirPath, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
