package api

import (
	"context"
	"errors"
	"hgnextfs/internal/entities"
	"hgnextfs/open_api/agentAPI"
	"io"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

type ExportUseCase interface {
	Create(ctx context.Context, data entities.ExportData) error
}

type FileUseCase interface {
	Create(ctx context.Context, fileID uuid.UUID, body io.Reader) error
	Delete(ctx context.Context, fileID uuid.UUID) error
	Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error)
	IDs(ctx context.Context) ([]uuid.UUID, error)
}

type Controller struct {
	startAt time.Time
	logger  *slog.Logger
	tracer  trace.Tracer
	addr    string
	debug   bool

	ogenServer *agentAPI.Server

	exportUseCase ExportUseCase
	fileUseCase   FileUseCase

	token string
}

func New(
	startAt time.Time,
	logger *slog.Logger,
	tracer trace.Tracer,
	exportUseCase ExportUseCase,
	fileUseCase FileUseCase,
	addr string,
	debug bool,
	token string,
) (*Controller, error) {
	c := &Controller{
		startAt:       startAt,
		logger:        logger,
		tracer:        tracer,
		addr:          addr,
		debug:         debug,
		token:         token,
		exportUseCase: exportUseCase,
		fileUseCase:   fileUseCase,
	}

	ogenServer, err := agentAPI.NewServer(c, c)
	if err != nil {
		return nil, err
	}

	c.ogenServer = ogenServer

	return c, nil
}

var errorAccessForbidden = errors.New("access forbidden")

func (c *Controller) HandleHeaderAuth(ctx context.Context, operationName string, t agentAPI.HeaderAuth) (context.Context, error) {
	if c.token == "" {
		return ctx, nil
	}

	if c.token != t.APIKey {
		return ctx, errorAccessForbidden
	}

	return ctx, nil
}
