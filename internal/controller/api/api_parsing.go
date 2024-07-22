package api

import (
	"context"
	"hgnextfs/internal/controller/api/internal/server"
)

func (c *Controller) APIParsingBookCheckPost(ctx context.Context, req *server.APIParsingBookCheckPostReq) (server.APIParsingBookCheckPostRes, error) {
	return &server.APIParsingBookCheckPostBadRequest{
		InnerCode: ValidationCode,
		Details:   server.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIParsingBookPost(ctx context.Context, req *server.APIParsingBookPostReq) (server.APIParsingBookPostRes, error) {
	return &server.APIParsingBookPostBadRequest{
		InnerCode: ValidationCode,
		Details:   server.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIParsingPageCheckPost(ctx context.Context, req *server.APIParsingPageCheckPostReq) (server.APIParsingPageCheckPostRes, error) {
	return &server.APIParsingPageCheckPostBadRequest{
		InnerCode: ValidationCode,
		Details:   server.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIParsingPagePost(ctx context.Context, req *server.APIParsingPagePostReq) (server.APIParsingPagePostRes, error) {
	return &server.APIParsingPagePostBadRequest{
		InnerCode: ValidationCode,
		Details:   server.NewOptString("unsupported api"),
	}, nil
}
