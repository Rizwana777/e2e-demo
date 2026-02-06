package app

import "sync"

type Store struct {
	mu    sync.Mutex
	Items []string
}

func NewStore() *Store {
	return &Store{
		Items: []string{},
	}
}

func (s *Store) Add(item string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Items = append(s.Items, item)
}

func (s *Store) List() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]string{}, s.Items...)
}
