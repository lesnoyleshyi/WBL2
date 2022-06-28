package output

import (
	"WBL2/Tasks/T11/internal/domain/entity"
	"context"
	"time"
)

type EventsStorage interface {
	Save(ctx context.Context, event entity.Event) error
	Update(ctx context.Context, event entity.Event) error
	Delete(ctx context.Context, event entity.Event) error
	Get(ctx context.Context, from, to time.Time) ([]entity.Event, error)
}
