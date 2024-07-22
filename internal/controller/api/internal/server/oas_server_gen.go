// Code generated by ogen, DO NOT EDIT.

package server

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// APICoreStatusGet implements GET /api/core/status operation.
	//
	// Получение данных о состоянии агента.
	//
	// GET /api/core/status
	APICoreStatusGet(ctx context.Context) (APICoreStatusGetRes, error)
	// APIExportArchivePost implements POST /api/export/archive operation.
	//
	// Загрузка архива.
	//
	// POST /api/export/archive
	APIExportArchivePost(ctx context.Context, req APIExportArchivePostReq, params APIExportArchivePostParams) (APIExportArchivePostRes, error)
	// APIFsCreatePost implements POST /api/fs/create operation.
	//
	// Создание нового файла.
	//
	// POST /api/fs/create
	APIFsCreatePost(ctx context.Context, req APIFsCreatePostReq, params APIFsCreatePostParams) (APIFsCreatePostRes, error)
	// APIFsDeletePost implements POST /api/fs/delete operation.
	//
	// Удаление файла.
	//
	// POST /api/fs/delete
	APIFsDeletePost(ctx context.Context, req *APIFsDeletePostReq) (APIFsDeletePostRes, error)
	// APIFsGetGet implements GET /api/fs/get operation.
	//
	// Получение файла.
	//
	// GET /api/fs/get
	APIFsGetGet(ctx context.Context, params APIFsGetGetParams) (APIFsGetGetRes, error)
	// APIFsIdsGet implements GET /api/fs/ids operation.
	//
	// Получение ID всех хранимых файлов.
	//
	// GET /api/fs/ids
	APIFsIdsGet(ctx context.Context) (APIFsIdsGetRes, error)
	// APIParsingBookCheckPost implements POST /api/parsing/book/check operation.
	//
	// Предварительная проверка ссылок на новые книги.
	//
	// POST /api/parsing/book/check
	APIParsingBookCheckPost(ctx context.Context, req *APIParsingBookCheckPostReq) (APIParsingBookCheckPostRes, error)
	// APIParsingBookPost implements POST /api/parsing/book operation.
	//
	// Обработка новой книги.
	//
	// POST /api/parsing/book
	APIParsingBookPost(ctx context.Context, req *APIParsingBookPostReq) (APIParsingBookPostRes, error)
	// APIParsingPageCheckPost implements POST /api/parsing/page/check operation.
	//
	// Предварительная проверка ссылок для загрузки страниц.
	//
	// POST /api/parsing/page/check
	APIParsingPageCheckPost(ctx context.Context, req *APIParsingPageCheckPostReq) (APIParsingPageCheckPostRes, error)
	// APIParsingPagePost implements POST /api/parsing/page operation.
	//
	// Загрузка изображения страницы.
	//
	// POST /api/parsing/page
	APIParsingPagePost(ctx context.Context, req *APIParsingPagePostReq) (APIParsingPagePostRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
