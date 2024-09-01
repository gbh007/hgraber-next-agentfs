// Code generated by ogen, DO NOT EDIT.

package serverAPI

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// APIAgentDeletePost implements POST /api/agent/delete operation.
//
// Удаление агента.
//
// POST /api/agent/delete
func (UnimplementedHandler) APIAgentDeletePost(ctx context.Context, req *APIAgentDeletePostReq) (r APIAgentDeletePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAgentListPost implements POST /api/agent/list operation.
//
// Список агентов.
//
// POST /api/agent/list
func (UnimplementedHandler) APIAgentListPost(ctx context.Context, req *APIAgentListPostReq) (r APIAgentListPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAgentNewPost implements POST /api/agent/new operation.
//
// Создание нового агента.
//
// POST /api/agent/new
func (UnimplementedHandler) APIAgentNewPost(ctx context.Context, req *APIAgentNewPostReq) (r APIAgentNewPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAgentTaskExportPost implements POST /api/agent/task/export operation.
//
// Экспорт книг в другую систему.
//
// POST /api/agent/task/export
func (UnimplementedHandler) APIAgentTaskExportPost(ctx context.Context, req *APIAgentTaskExportPostReq) (r APIAgentTaskExportPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookArchiveIDGet implements GET /api/book/archive/{id} operation.
//
// Получение архива с книгой.
//
// GET /api/book/archive/{id}
func (UnimplementedHandler) APIBookArchiveIDGet(ctx context.Context, params APIBookArchiveIDGetParams) (r APIBookArchiveIDGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookDeletePost implements POST /api/book/delete operation.
//
// Удаление книги (без удаления метаинформации).
//
// POST /api/book/delete
func (UnimplementedHandler) APIBookDeletePost(ctx context.Context, req *APIBookDeletePostReq) (r APIBookDeletePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookDetailsPost implements POST /api/book/details operation.
//
// Информация о книге.
//
// POST /api/book/details
func (UnimplementedHandler) APIBookDetailsPost(ctx context.Context, req *APIBookDetailsPostReq) (r APIBookDetailsPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookListPost implements POST /api/book/list operation.
//
// Список книг.
//
// POST /api/book/list
func (UnimplementedHandler) APIBookListPost(ctx context.Context, req *APIBookListPostReq) (r APIBookListPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookRawPost implements POST /api/book/raw operation.
//
// Информация о книге (или по ИД или по адресу).
//
// POST /api/book/raw
func (UnimplementedHandler) APIBookRawPost(ctx context.Context, req *APIBookRawPostReq) (r APIBookRawPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookVerifyPost implements POST /api/book/verify operation.
//
// Подтверждение (модерация) книги, нужна в случае
// массовой обработки.
//
// POST /api/book/verify
func (UnimplementedHandler) APIBookVerifyPost(ctx context.Context, req *APIBookVerifyPostReq) (r APIBookVerifyPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIFileIDGet implements GET /api/file/{id} operation.
//
// Получение тела файла (изображения страницы).
//
// GET /api/file/{id}
func (UnimplementedHandler) APIFileIDGet(ctx context.Context, params APIFileIDGetParams) (r APIFileIDGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIPageBodyPost implements POST /api/page/body operation.
//
// Получение тела страницы (по оригинальному адресу или
// данным книги).
//
// POST /api/page/body
func (UnimplementedHandler) APIPageBodyPost(ctx context.Context, req *APIPageBodyPostReq) (r APIPageBodyPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIParsingBookExistsPost implements POST /api/parsing/book/exists operation.
//
// Проверка наличия ссылок на книги.
//
// POST /api/parsing/book/exists
func (UnimplementedHandler) APIParsingBookExistsPost(ctx context.Context, req *APIParsingBookExistsPostReq) (r APIParsingBookExistsPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIParsingPageExistsPost implements POST /api/parsing/page/exists operation.
//
// Проверка наличия ссылок для страниц.
//
// POST /api/parsing/page/exists
func (UnimplementedHandler) APIParsingPageExistsPost(ctx context.Context, req *APIParsingPageExistsPostReq) (r APIParsingPageExistsPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APISystemDeduplicateArchivePost implements POST /api/system/deduplicate/archive operation.
//
// Проверка наличия данных в системе из архива.
//
// POST /api/system/deduplicate/archive
func (UnimplementedHandler) APISystemDeduplicateArchivePost(ctx context.Context, req APISystemDeduplicateArchivePostReq) (r APISystemDeduplicateArchivePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APISystemHandlePost implements POST /api/system/handle operation.
//
// Обработка ссылок на новые книги.
//
// POST /api/system/handle
func (UnimplementedHandler) APISystemHandlePost(ctx context.Context, req *APISystemHandlePostReq) (r APISystemHandlePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APISystemImportArchivePost implements POST /api/system/import/archive operation.
//
// Импорт новой книги через архив.
//
// POST /api/system/import/archive
func (UnimplementedHandler) APISystemImportArchivePost(ctx context.Context, req APISystemImportArchivePostReq) (r APISystemImportArchivePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APISystemInfoGet implements GET /api/system/info operation.
//
// Получение общей информации о системе.
//
// GET /api/system/info
func (UnimplementedHandler) APISystemInfoGet(ctx context.Context) (r APISystemInfoGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APISystemRPCDeduplicateFilesPost implements POST /api/system/rpc/deduplicate/files operation.
//
// Дедупликация файлов.
//
// POST /api/system/rpc/deduplicate/files
func (UnimplementedHandler) APISystemRPCDeduplicateFilesPost(ctx context.Context) (r APISystemRPCDeduplicateFilesPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APISystemRPCRemoveDetachedFilesPost implements POST /api/system/rpc/remove/detached-files operation.
//
// Удаление несвязанных файлов.
//
// POST /api/system/rpc/remove/detached-files
func (UnimplementedHandler) APISystemRPCRemoveDetachedFilesPost(ctx context.Context) (r APISystemRPCRemoveDetachedFilesPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APISystemRPCRemoveMismatchFilesPost implements POST /api/system/rpc/remove/mismatch-files operation.
//
// Удаление рассинхронизированных файлов
// (несоответствие файловой системы и БД).
//
// POST /api/system/rpc/remove/mismatch-files
func (UnimplementedHandler) APISystemRPCRemoveMismatchFilesPost(ctx context.Context) (r APISystemRPCRemoveMismatchFilesPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APISystemWorkerConfigPost implements POST /api/system/worker/config operation.
//
// Динамическая конфигурация раннеров (воркеров),
// сбрасывается при перезапуске системы.
//
// POST /api/system/worker/config
func (UnimplementedHandler) APISystemWorkerConfigPost(ctx context.Context, req *APISystemWorkerConfigPostReq) (r APISystemWorkerConfigPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIUserLoginPost implements POST /api/user/login operation.
//
// Проставление токена в куки.
//
// POST /api/user/login
func (UnimplementedHandler) APIUserLoginPost(ctx context.Context, req *APIUserLoginPostReq) (r APIUserLoginPostRes, _ error) {
	return r, ht.ErrNotImplemented
}
