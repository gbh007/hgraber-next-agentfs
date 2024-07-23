package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// FIXME: реализовать
func (c *Controller) Serve(parentCtx context.Context) error {
	server := &http.Server{
		Handler: otelPropagation(c.logIO(cors(c.ogenServer))),
		Addr:    c.addr,
	}

	go func() {
		<-parentCtx.Done()
		c.logger.InfoContext(parentCtx, "stopping api server")

		shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(parentCtx), time.Second*10)
		defer cancel()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			c.logger.ErrorContext(parentCtx, err.Error())
		}
	}()

	c.logger.InfoContext(parentCtx, "api server start")
	defer c.logger.InfoContext(parentCtx, "api server stop")

	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)

			return
		}

		if next != nil {
			next.ServeHTTP(w, r)
		}
	})
}

func otelPropagation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := otel.GetTextMapPropagator().Extract(
			r.Context(),
			propagation.HeaderCarrier(r.Header),
		)

		if next != nil {
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
