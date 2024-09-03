package api

import (
	"context"
	"errors"
	"hgnextfs/internal/entities"
	"hgnextfs/open_api/agentAPI"
)

func (c *Controller) APIFsCreatePost(ctx context.Context, req agentAPI.APIFsCreatePostReq, params agentAPI.APIFsCreatePostParams) (agentAPI.APIFsCreatePostRes, error) {
	if c.fileUseCase == nil {
		return &agentAPI.APIFsCreatePostBadRequest{
			InnerCode: ValidationCode,
			Details:   agentAPI.NewOptString("unsupported api"),
		}, nil
	}

	err := c.fileUseCase.Create(ctx, params.FileID, req.Data)
	if errors.Is(err, entities.FileAlreadyExistsError) {
		return &agentAPI.APIFsCreatePostConflict{
			InnerCode: FileUseCaseCode,
			Details:   agentAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &agentAPI.APIFsCreatePostInternalServerError{
			InnerCode: FileUseCaseCode,
			Details:   agentAPI.NewOptString(err.Error()),
		}, nil
	}

	return &agentAPI.APIFsCreatePostNoContent{}, nil
}

func (c *Controller) APIFsDeletePost(ctx context.Context, req *agentAPI.APIFsDeletePostReq) (agentAPI.APIFsDeletePostRes, error) {
	if c.fileUseCase == nil {
		return &agentAPI.APIFsDeletePostBadRequest{
			InnerCode: ValidationCode,
			Details:   agentAPI.NewOptString("unsupported api"),
		}, nil
	}

	err := c.fileUseCase.Delete(ctx, req.FileID)
	if errors.Is(err, entities.FileNotFoundError) {
		return &agentAPI.APIFsDeletePostNotFound{
			InnerCode: FileUseCaseCode,
			Details:   agentAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &agentAPI.APIFsDeletePostInternalServerError{
			InnerCode: FileUseCaseCode,
			Details:   agentAPI.NewOptString(err.Error()),
		}, nil
	}

	return &agentAPI.APIFsDeletePostNoContent{}, nil
}

func (c *Controller) APIFsGetGet(ctx context.Context, params agentAPI.APIFsGetGetParams) (agentAPI.APIFsGetGetRes, error) {
	if c.fileUseCase == nil {
		return &agentAPI.APIFsGetGetBadRequest{
			InnerCode: ValidationCode,
			Details:   agentAPI.NewOptString("unsupported api"),
		}, nil
	}

	body, err := c.fileUseCase.Get(ctx, params.FileID)
	if errors.Is(err, entities.FileNotFoundError) {
		return &agentAPI.APIFsGetGetNotFound{
			InnerCode: FileUseCaseCode,
			Details:   agentAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &agentAPI.APIFsGetGetInternalServerError{
			InnerCode: FileUseCaseCode,
			Details:   agentAPI.NewOptString(err.Error()),
		}, nil
	}

	// FIXME: работать с типом контента как в основном сервере
	return &agentAPI.APIFsGetGetOK{
		Data: body,
	}, nil
}

func (c *Controller) APIFsIdsGet(ctx context.Context) (agentAPI.APIFsIdsGetRes, error) {
	if c.fileUseCase == nil {
		return &agentAPI.APIFsIdsGetBadRequest{
			InnerCode: ValidationCode,
			Details:   agentAPI.NewOptString("unsupported api"),
		}, nil
	}

	ids, err := c.fileUseCase.IDs(ctx)
	if err != nil {
		return &agentAPI.APIFsIdsGetInternalServerError{
			InnerCode: FileUseCaseCode,
			Details:   agentAPI.NewOptString(err.Error()),
		}, nil
	}

	resp := agentAPI.APIFsIdsGetOKApplicationJSON(ids)

	return &resp, nil
}
