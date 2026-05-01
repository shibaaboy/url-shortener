package storage

import "sync"

type Storage struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]string),
	}
}

func (s *Storage) Get(id string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[id]
	return val, ok
}

func (s *Storage) Set(id, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[id] = value
}
