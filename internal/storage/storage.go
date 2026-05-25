package storage

import (
	"encoding/json"
	"os"
	"sync"
)

type File struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

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

func (s *Storage) toFileFormat() []File {
	var list []File
	for short, original := range s.data {
		list = append(list, File{
			UUID:        "",
			ShortURL:    short,
			OriginalURL: original,
		})
	}

	return list
}

func (s *Storage) SaveToFile(path string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	if err != nil {
		return
	}

	defer file.Close()
	json.NewEncoder(file).Encode(s.toFileFormat())

}

func (s *Storage) LoadFromFile(path string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	file, err := os.OpenFile(path, os.O_RDONLY, 0664)

	if err != nil {
		return
	}

	defer file.Close()

	var f []File
	json.NewDecoder(file).Decode(&f)

	for _, item := range f {
		s.data[item.ShortURL] = item.OriginalURL
	}
}
