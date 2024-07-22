package api

import (
	"context"
	"hgnextfs/internal/controller/api/internal/server"
)

func (c *Controller) APICoreStatusGet(ctx context.Context) (server.APICoreStatusGetRes, error) {
	return &server.APICoreStatusGetOK{
		StartAt: c.startAt,
		Status:  server.APICoreStatusGetOKStatusOk,
	}, nil
}
