package http

import (
	"WBL2/Tasks/T11/internal/domain/entity"
	"context"
	"github.com/mailru/easyjson"
	"net/http"
	"time"
)

func (s Server) eventsHandler() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/create_event", s.createEvent)
	r.HandleFunc("/update_event", s.updateEvent)
	r.HandleFunc("/delete_event", s.deleteEvent)

	return r
}

func (s Server) createEvent(w http.ResponseWriter, r *http.Request) {
	var event entity.Event

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*respTimeout)
	defer cancel()

	err := easyjson.UnmarshalFromReader(r.Body, &event)
	defer func() { _ = r.Body.Close() }()
	if err != nil {
		s.logger.Info("marshalling error")
		return
	}

	if err := s.events.Create(ctx, event); err != nil {
		s.logger.Info(err)
		return
	}
}

func (s Server) updateEvent(w http.ResponseWriter, r *http.Request) {

}

func (s Server) deleteEvent(w http.ResponseWriter, r *http.Request) {

}
