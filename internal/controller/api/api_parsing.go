package api

import (
	"context"
	"hgnextfs/open_api/agentAPI"
)

func (c *Controller) APIParsingBookCheckPost(ctx context.Context, req *agentAPI.APIParsingBookCheckPostReq) (agentAPI.APIParsingBookCheckPostRes, error) {
	return &agentAPI.APIParsingBookCheckPostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentAPI.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIParsingBookPost(ctx context.Context, req *agentAPI.APIParsingBookPostReq) (agentAPI.APIParsingBookPostRes, error) {
	return &agentAPI.APIParsingBookPostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentAPI.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIParsingPageCheckPost(ctx context.Context, req *agentAPI.APIParsingPageCheckPostReq) (agentAPI.APIParsingPageCheckPostRes, error) {
	return &agentAPI.APIParsingPageCheckPostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentAPI.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIParsingPagePost(ctx context.Context, req *agentAPI.APIParsingPagePostReq) (agentAPI.APIParsingPagePostRes, error) {
	return &agentAPI.APIParsingPagePostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentAPI.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIParsingBookMultiPost(ctx context.Context, req *agentAPI.APIParsingBookMultiPostReq) (agentAPI.APIParsingBookMultiPostRes, error) {
	return &agentAPI.APIParsingBookMultiPostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentAPI.NewOptString("unsupported api"),
	}, nil
}
