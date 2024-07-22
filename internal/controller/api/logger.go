package api

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
)

func (c *Controller) logIO(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !c.debug {
			if next != nil {
				next.ServeHTTP(w, r)
			}

			return
		}

		var (
			responseData = "ignoring"
			requestData  = "ignoring"
		)

		if r.URL.Path != "/api/export/archive" && r.URL.Path != "/api/fs/create" {
			requestDataRaw, err := io.ReadAll(r.Body)
			if err != nil {
				c.logger.ErrorContext(
					r.Context(), "read request to log",
					slog.Any("error", err),
				)
			}

			requestData = string(requestDataRaw)

			r.Body.Close()
			r.Body = io.NopCloser(bytes.NewReader(requestDataRaw))
		}

		rw := newResponseWrapper(w)

		if next != nil {
			next.ServeHTTP(rw, r)
		}

		if r.URL.Path != "/api/parsing/page" && r.URL.Path != "/api/fs/get" {
			responseData = rw.body.String()
		}

		c.logger.DebugContext(
			r.Context(), "http request",
			slog.String("path", r.URL.Path),
			slog.String("method", r.Method),
			slog.Group(
				"request",
				slog.Any("headers", r.Header),
				slog.String("body", requestData),
			),
			slog.Group(
				"response",
				slog.Int("code", rw.statusCode),
				slog.Any("headers", rw.origin.Header()),
				slog.String("body", responseData),
			),
		)
	})
}

type responseWrapper struct {
	origin http.ResponseWriter

	statusCode int
	body       *bytes.Buffer
}

func newResponseWrapper(origin http.ResponseWriter) *responseWrapper {
	return &responseWrapper{
		origin: origin,
		body:   &bytes.Buffer{},
	}
}

func (rw *responseWrapper) Header() http.Header {
	return rw.origin.Header()
}

func (rw *responseWrapper) Write(data []byte) (int, error) {
	_, _ = rw.body.Write(data)

	return rw.origin.Write(data)
}

func (rw *responseWrapper) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.origin.WriteHeader(statusCode)
}
