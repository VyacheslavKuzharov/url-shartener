package infile

import (
	"context"
)

func (s *FileStorage) GetURL(ctx context.Context, key string) (string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return "", nil
}
