package http

import (
	"WBL2/Tasks/T11/internal/domain/entity"
	"context"
	"github.com/mailru/easyjson"
	"net/http"
	"time"
)

func (a AdapterHTTP) eventsHandler() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/create_event", a.createEvent)
	r.HandleFunc("/update_event", a.updateEvent)
	r.HandleFunc("/delete_event", a.deleteEvent)

	r.HandleFunc("/events_for_day", a.getEventsByPeriod)

	return r
}

func (a AdapterHTTP) createEvent(w http.ResponseWriter, r *http.Request) {
	var event entity.Event

	if r.Method != http.MethodPost {
		http.Error(w, "{\"error\":\"wrong method\"}", http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*respTimeout)
	defer cancel()

	err := easyjson.UnmarshalFromReader(r.Body, &event)
	defer func() { _ = r.Body.Close() }()
	if err != nil {
		a.logger.Info("marshalling error")
		http.Error(w, "{\"error\":\"can't read body\"}", http.StatusBadRequest)
		return
	}

	if err := a.events.Create(ctx, event); err != nil {
		a.logger.Info(err)
		http.Error(w, "{\"error\":\"can't create event\"}", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("{\"result\":\"event created\"}"))
}

func (a AdapterHTTP) updateEvent(w http.ResponseWriter, r *http.Request) {
	var event entity.Event

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if r.Method != http.MethodPost {
		http.Error(w, "{\"error\":\"wrong method\"}", http.StatusMethodNotAllowed)
		return
	}

	err := easyjson.UnmarshalFromReader(r.Body, &event)
	defer func() { _ = r.Body.Close() }()
	if err != nil {
		a.logger.Info("marshalling error")
		http.Error(w, "{\"error\":\"can't read body\"}", http.StatusBadRequest)
		return
	}

	if err := a.events.Update(ctx, event); err != nil {
		a.logger.Info(err)
		http.Error(w, "{\"error\":\"can't update event\"}", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("{\"result\":\"event updated\"}"))
}

func (a AdapterHTTP) deleteEvent(w http.ResponseWriter, r *http.Request) {
	var event entity.Event

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if r.Method != http.MethodPost {
		http.Error(w, "{\"error\":\"wrong method\"}", http.StatusMethodNotAllowed)
		return
	}

	// TODO Можно заменить на функцию, принирмающую reader, а возвращающую Event
	// пускай внутри сразу закрывает reader
	err := easyjson.UnmarshalFromReader(r.Body, &event)
	defer func() { _ = r.Body.Close() }()
	if err != nil {
		a.logger.Info("marshalling error")
		http.Error(w, "{\"error\":\"can't read body\"}", http.StatusBadRequest)
		return
	}

	if err := a.events.Delete(ctx, event); err != nil {
		// TODO Можно заменить две нижние строки на respondError
		a.logger.Info(err)
		http.Error(w, "{\"error\":\"can't delete event\"}", http.StatusInternalServerError)

		return
	}

	// TODO Можно заменить две нижние строки на respondSuccess
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("{\"result\":\"event deleted\"}"))
}

func (a AdapterHTTP) getEventsByPeriod(w http.ResponseWriter, r *http.Request) {
	var events []entity.Event
	var err error

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if r.Method != http.MethodGet {
		http.Error(w, "{\"error\":\"wrong method\"}", http.StatusMethodNotAllowed)
		return
	}

	events, err = a.events.GetByPeriod(ctx, "TODO")
	if err != nil {
		a.logger.Info(err)
		http.Error(w, "{\"error\":\"can't retrieve events\"}", http.StatusInternalServerError)
	}

	for _, event := range events {
		_, _ = w.Write([]byte(event.String()))
	}
}
