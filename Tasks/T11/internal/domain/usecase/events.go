package usecase

import (
	"WBL2/Tasks/T11/internal/domain/entity"
	"WBL2/Tasks/T11/internal/ports"
	"context"
)

type Service struct {
	db ports.EventsStorage
}

func New(db ports.EventsStorage) Service {
	return Service{
		db: db,
	}
}

func (s Service) CreateEvent(event entity.Event) error {
	return s.db.Create(context.TODO(), event)
}

func (s Service) UpdateEvent(event entity.Event) error {
	return s.db.Update(context.TODO(), event)
}

func (s Service) DeleteEvent(event entity.Event) error {
	return s.db.Delete(context.TODO(), event)
}

func (s Service) GetEventsByPeriod(period string) ([]entity.Event, error) {
	return s.db.GetByPeriod(context.TODO(), period)
}
