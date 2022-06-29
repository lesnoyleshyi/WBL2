package http

import (
	"WBL2/Tasks/T11/internal/ports/input"
	"context"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type AdapterHTTP struct {
	server *http.Server
	events input.EventsService
	logger *zap.SugaredLogger
}

const respTimeout = 5

func New(events input.EventsService, logger *zap.SugaredLogger) AdapterHTTP {
	var a AdapterHTTP

	a.events = events
	a.logger = logger
	a.server = &http.Server{
		Handler: a.loggerWrapper(a.routes()),
	}

	return a
}

func (a AdapterHTTP) routes() http.Handler {
	r := http.NewServeMux()

	r.Handle("/", a.eventsHandler())

	return r
}

func (a AdapterHTTP) Start(ctx context.Context) error {
	err := a.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (a AdapterHTTP) Stop(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}

func (a AdapterHTTP) loggerWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		url := r.RequestURI
		ip := r.RemoteAddr

		next.ServeHTTP(w, r)

		finish := time.Now()
		serveTime := finish.Sub(start)
		a.logger.Infow("request served",
			zap.String("url", url),
			zap.String("from", ip),
			zap.Duration("resp time", serveTime),
		)
	})
}
