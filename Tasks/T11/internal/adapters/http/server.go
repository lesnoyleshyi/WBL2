package http

import (
	"WBL2/Tasks/T11/internal/ports"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	server *http.Server
	events ports.EventsStorage
	logger *zap.SugaredLogger
}

const respTimeout = 5

func New(events ports.EventsStorage, logger *zap.SugaredLogger) Server {
	var s Server

	s.events = events
	s.logger = logger
	s.server.Handler = s.routes()

	return s
}

func (s Server) routes() http.Handler {
	r := http.NewServeMux()

	r.Handle("/", s.eventsHandler())

	return r
}

func (s Server) respondJSON(w http.ResponseWriter, status int) {

}

func (s Server) respondError() {

}
