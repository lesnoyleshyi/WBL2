package http

import (
	"WBL2/Tasks/T11/internal/adapters/http/utils"
	"WBL2/Tasks/T11/internal/domain/entity"
	"context"
	"net/http"
	"time"
)

func (a AdapterHTTP) eventsHandler() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/create_event", a.createEvent)
	r.HandleFunc("/update_event", a.updateEvent)
	r.HandleFunc("/delete_event", a.deleteEvent)

	r.HandleFunc("/events_for_day", a.getEventsByPeriod)
	r.HandleFunc("/events_for_week", a.getEventsByPeriod)
	r.HandleFunc("/events_for_month", a.getEventsByPeriod)

	return r
}

func (a AdapterHTTP) createEvent(w http.ResponseWriter, r *http.Request) {
	// TODO Этому место в миддлваре
	if r.Method != http.MethodPost {
		http.Error(w, "{\"error\":\"wrong method\"}", http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*respTimeout)
	defer cancel()

	event, err := utils.GetEvent(r.Body)
	if err != nil {
		a.respondError(w, "can't read body", http.StatusBadRequest, err)
		return
	}

	if err := a.events.Create(ctx, *event); err != nil {
		a.respondError(w, "can't create event", http.StatusInternalServerError, err)
		return
	}

	a.respondSuccess(w, "event created", http.StatusCreated)
}

func (a AdapterHTTP) updateEvent(w http.ResponseWriter, r *http.Request) {
	// TODO Этому место в миддлваре
	if r.Method != http.MethodPost {
		http.Error(w, "{\"error\":\"wrong method\"}", http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	event, err := utils.GetEvent(r.Body)
	if err != nil {
		a.respondError(w, "can't read body", http.StatusBadRequest, err)
		return
	}

	if err := a.events.Update(ctx, *event); err != nil {
		a.respondError(w, "can't update event", http.StatusInternalServerError, err)
		return
	}

	a.respondSuccess(w, "event updated", http.StatusOK)
}

func (a AdapterHTTP) deleteEvent(w http.ResponseWriter, r *http.Request) {
	// TODO Этому место в миддлваре
	if r.Method != http.MethodPost {
		http.Error(w, "{\"error\":\"wrong method\"}", http.StatusMethodNotAllowed)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	event, err := utils.GetEvent(r.Body)
	if err != nil {
		a.respondError(w, "can't read body", http.StatusBadRequest, err)
		return
	}

	if err := a.events.Delete(ctx, *event); err != nil {
		a.respondError(w, "can't delete event", http.StatusInternalServerError, err)
		return
	}

	a.respondSuccess(w, "event deleted", http.StatusOK)
}

func (a AdapterHTTP) getEventsByPeriod(w http.ResponseWriter, r *http.Request) {
	// TODO Этому место в миддлваре
	if r.Method != http.MethodGet {
		http.Error(w, "{\"error\":\"wrong method\"}", http.StatusMethodNotAllowed)
		return
	}
	var events []entity.Event
	var err error

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	from, to, err := utils.GetPeriod(r)
	if err != nil {
		a.respondError(w, "can't recognise query parameters", http.StatusBadRequest, err)
		return
	}

	events, err = a.events.Get(ctx, from, to)
	if err != nil {
		a.respondError(w, "can't retrieve events", http.StatusInternalServerError, err)
		return
	}

	//TODO Сделать нормальный ответ
	for _, event := range events {
		_, _ = w.Write([]byte(event.String()))
	}
	a.respondSuccess(w, utils.ArrayPresenter(events), http.StatusOK)
}
