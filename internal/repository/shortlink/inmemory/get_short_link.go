package inmemory

import (
	"context"
	"errors"
)

func (s *MemStorage) GetURL(ctx context.Context, key string) (string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	originalURL, ok := s.storage[key]
	if !ok {
		return "", errors.New("shortKey not found")
	}

	return originalURL, nil
}
