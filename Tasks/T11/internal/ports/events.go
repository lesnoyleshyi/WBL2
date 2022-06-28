package ports

import (
	"WBL2/Tasks/T11/internal/domain/entity"
	"context"
)

type EventsStorage interface {
	Create(ctx context.Context, event entity.Event) error
	Update(ctx context.Context, event entity.Event) error
	Delete(ctx context.Context, event entity.Event) error
	GetByPeriod(ctx context.Context, period string) ([]entity.Event, error)
}
