package usecase

import (
	"WBL2/Tasks/T11/internal/domain/entity"
	"WBL2/Tasks/T11/internal/ports/output"
	"context"
	"errors"
	"time"
)

type Service struct {
	db output.EventsStorage
}

func New(db output.EventsStorage) Service {
	return Service{
		db: db,
	}
}

func (s Service) Create(ctx context.Context, event entity.Event) error {
	return s.db.Save(ctx, event)
}

func (s Service) Update(ctx context.Context, event entity.Event) error {
	return s.db.Update(ctx, event)
}

func (s Service) Delete(ctx context.Context, event entity.Event) error {
	return s.db.Delete(ctx, event)
}

func (s Service) GetByPeriod(ctx context.Context, period string) ([]entity.Event, error) {
	if period == "" {
		return nil, errors.New("shlyapa")
	}

	from := time.Now()
	to := time.Now().Add(time.Hour)
	return s.db.Get(ctx, from, to)
}
