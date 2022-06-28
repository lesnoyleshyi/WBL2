package http

import (
	"WBL2/Tasks/T11/internal/ports/input"
	"context"
	"errors"
	"go.uber.org/zap"
	"net/http"
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
		Handler: a.routes(),
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

func (a AdapterHTTP) respondSuccess(w http.ResponseWriter, status int) {

}

func (a AdapterHTTP) respondError() {

}
