package api

import (
	"context"
	"errors"
	"hgnextfs/internal/controller/api/internal/server"
	"io"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type exportUseCase interface {
	Create(ctx context.Context, bookID uuid.UUID, bookName string, body io.Reader) error
}

type fileUseCase interface {
	Create(ctx context.Context, fileID uuid.UUID, body io.Reader) error
	Delete(ctx context.Context, fileID uuid.UUID) error
	Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error)
	IDs(ctx context.Context) ([]uuid.UUID, error)
}

type Controller struct {
	startAt time.Time
	logger  *slog.Logger
	addr    string
	debug   bool

	ogenServer *server.Server

	exportUseCase exportUseCase
	fileUseCase   fileUseCase

	token string
}

func New(
	startAt time.Time,
	logger *slog.Logger,
	exportUseCase exportUseCase,
	fileUseCase fileUseCase,
	addr string,
	debug bool,
	token string,
) (*Controller, error) {
	c := &Controller{
		startAt:       startAt,
		logger:        logger,
		addr:          addr,
		debug:         debug,
		token:         token,
		exportUseCase: exportUseCase,
		fileUseCase:   fileUseCase,
	}

	ogenServer, err := server.NewServer(c, c)
	if err != nil {
		return nil, err
	}

	c.ogenServer = ogenServer

	return c, nil
}

var errorAccessForbidden = errors.New("access forbidden")

func (c *Controller) HandleHeaderAuth(ctx context.Context, operationName string, t server.HeaderAuth) (context.Context, error) {
	if c.token == "" {
		return ctx, nil
	}

	if c.token != t.APIKey {
		return ctx, errorAccessForbidden
	}

	return ctx, nil
}

func (c *Controller) APIFsCreatePost(ctx context.Context, req server.APIFsCreatePostReq, params server.APIFsCreatePostParams) (server.APIFsCreatePostRes, error) {
	return nil, nil
}

func (c *Controller) APIFsDeletePost(ctx context.Context, req *server.APIFsDeletePostReq) (server.APIFsDeletePostRes, error) {
	return nil, nil
}

func (c *Controller) APIFsGetGet(ctx context.Context, params server.APIFsGetGetParams) (server.APIFsGetGetRes, error) {
	return nil, nil
}

func (c *Controller) APIFsIdsGet(ctx context.Context) (server.APIFsIdsGetRes, error) {
	return nil, nil
}
