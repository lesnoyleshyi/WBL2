package cache

import (
	"WBL2/Tasks/T11/internal/domain/entity"
	"context"
	"time"
)

func (c Cache) Save(ctx context.Context, event entity.Event) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		c.mu.Lock()
		c.storage[event.Title] = &event
		c.mu.Unlock()

		return nil
	}
}

func (c Cache) Update(ctx context.Context, event entity.Event) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		c.mu.Lock()
		c.storage[event.Title] = &event
		c.mu.Unlock()

		return nil
	}
}

func (c Cache) Delete(ctx context.Context, event entity.Event) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		c.mu.Lock()
		c.storage[event.Title] = nil
		c.mu.Unlock()

		return nil
	}
}

func (c Cache) Get(ctx context.Context, from, to time.Time) ([]entity.Event, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		var events []entity.Event

		c.mu.RLock()

		for _, event := range c.storage {
			if fit(event.Datetime, from, to) {
				events = append(events, *event)
			}
		}

		c.mu.RUnlock()

		return events, nil
	}
}

func fit(datetime, from, to time.Time) bool {
	return datetime.After(from) && datetime.Before(to)
}
