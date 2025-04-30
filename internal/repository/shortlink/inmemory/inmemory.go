package inmemory

import (
	"sync"
)

type MemStorage struct {
	mutex   sync.RWMutex
	storage map[string]string
}

func NewStorage() *MemStorage {
	return &MemStorage{
		storage: make(map[string]string),
	}
}
