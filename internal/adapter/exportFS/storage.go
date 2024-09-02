package exportFS

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"sync"
)

type Storage struct {
	fsPath string

	limitOnFolder int
	fsLimitMutex  sync.Mutex

	logger *slog.Logger
}

func New(path string, logger *slog.Logger, limitOnFolder int) (*Storage, error) {
	err := createDir(path)
	if err != nil {
		return nil, err
	}

	return &Storage{
		fsPath:        path,
		logger:        logger,
		limitOnFolder: limitOnFolder,
	}, nil
}

func createDir(dirPath string) error {
	info, err := os.Stat(dirPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if info != nil && info.IsDir() {
		return nil
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
