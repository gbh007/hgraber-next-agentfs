package api

import (
	"context"
	"errors"
	"hgnextfs/internal/controller/api/internal/server"
	"hgnextfs/internal/entities"
)

func (c *Controller) APIFsCreatePost(ctx context.Context, req server.APIFsCreatePostReq, params server.APIFsCreatePostParams) (server.APIFsCreatePostRes, error) {
	if c.fileUseCase == nil {
		return &server.APIFsCreatePostBadRequest{
			InnerCode: ValidationCode,
			Details:   server.NewOptString("unsupported api"),
		}, nil
	}

	err := c.fileUseCase.Create(ctx, params.FileID, req.Data)
	if errors.Is(err, entities.FileAlreadyExistsError) {
		return &server.APIFsCreatePostConflict{
			InnerCode: FileUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &server.APIFsCreatePostInternalServerError{
			InnerCode: FileUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.APIFsCreatePostNoContent{}, nil
}

func (c *Controller) APIFsDeletePost(ctx context.Context, req *server.APIFsDeletePostReq) (server.APIFsDeletePostRes, error) {
	if c.fileUseCase == nil {
		return &server.APIFsDeletePostBadRequest{
			InnerCode: ValidationCode,
			Details:   server.NewOptString("unsupported api"),
		}, nil
	}

	err := c.fileUseCase.Delete(ctx, req.FileID)
	if errors.Is(err, entities.FileNotFoundError) {
		return &server.APIFsDeletePostNotFound{
			InnerCode: FileUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &server.APIFsDeletePostInternalServerError{
			InnerCode: FileUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.APIFsDeletePostNoContent{}, nil
}

func (c *Controller) APIFsGetGet(ctx context.Context, params server.APIFsGetGetParams) (server.APIFsGetGetRes, error) {
	if c.fileUseCase == nil {
		return &server.APIFsGetGetBadRequest{
			InnerCode: ValidationCode,
			Details:   server.NewOptString("unsupported api"),
		}, nil
	}

	body, err := c.fileUseCase.Get(ctx, params.FileID)
	if errors.Is(err, entities.FileNotFoundError) {
		return &server.APIFsGetGetNotFound{
			InnerCode: FileUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &server.APIFsGetGetInternalServerError{
			InnerCode: FileUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	// FIXME: работать с типом контента как в основном сервере
	return &server.APIFsGetGetOK{
		Data: body,
	}, nil
}

func (c *Controller) APIFsIdsGet(ctx context.Context) (server.APIFsIdsGetRes, error) {
	if c.fileUseCase == nil {
		return &server.APIFsIdsGetBadRequest{
			InnerCode: ValidationCode,
			Details:   server.NewOptString("unsupported api"),
		}, nil
	}

	ids, err := c.fileUseCase.IDs(ctx)
	if err != nil {
		return &server.APIFsIdsGetInternalServerError{
			InnerCode: FileUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	resp := server.APIFsIdsGetOKApplicationJSON(ids)

	return &resp, nil
}
