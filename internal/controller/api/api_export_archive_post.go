package api

import (
	"context"
	"hgnextfs/internal/controller/api/internal/server"
)

func (c *Controller) APIExportArchivePost(ctx context.Context, req server.APIExportArchivePostReq, params server.APIExportArchivePostParams) (server.APIExportArchivePostRes, error) {
	if c.exportUseCase == nil {
		return &server.APIExportArchivePostBadRequest{
			InnerCode: ValidationCode,
			Details:   server.NewOptString("unsupported api"),
		}, nil
	}

	err := c.exportUseCase.Create(ctx, params.BookID, params.BookName, req.Data)
	if err != nil {
		return &server.APIExportArchivePostInternalServerError{
			InnerCode: ExportUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.APIExportArchivePostNoContent{}, nil
}
