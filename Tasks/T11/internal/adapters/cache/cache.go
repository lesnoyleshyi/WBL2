package cache

import (
	"WBL2/Tasks/T11/internal/domain/entity"
	"sync"
)

type Cache struct {
	mu      *sync.RWMutex
	storage map[string]*entity.Event
}

func New() Cache {
	RWMutex := new(sync.RWMutex)
	m := make(map[string]*entity.Event)

	return Cache{
		mu:      RWMutex,
		storage: m,
	}
}
