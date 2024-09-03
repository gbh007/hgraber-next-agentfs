package api

import (
	"context"
	"hgnextfs/open_api/agentAPI"
)

func (c *Controller) APICoreStatusGet(ctx context.Context) (agentAPI.APICoreStatusGetRes, error) {
	return &agentAPI.APICoreStatusGetOK{
		StartAt: c.startAt,
		Status:  agentAPI.APICoreStatusGetOKStatusOk,
	}, nil
}
