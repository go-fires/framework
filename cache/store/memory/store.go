package memory

import (
	"github.com/go-fires/fires/cache"
	"sync"
	"time"
)

type item struct {
	value   interface{}
	expired time.Time
}

func (i *item) isExpired() bool {
	return !i.expired.IsZero() && !i.expired.After(time.Now())
}

type Store struct {
	items map[string]*item
	rw    sync.RWMutex
}

var _ cache.Store = (*Store)(nil)

func New() *Store {
	return &Store{
		items: make(map[string]*item),
	}
}

func (s *Store) Has(key string) bool {
	s.rw.RLock()
	defer s.rw.RUnlock()

	if v, ok := s.items[key]; ok {
		if !v.isExpired() {
			return true
		} else {
			delete(s.items, key)
		}
	}

	return false
}

func (s *Store) Get(key string, dest interface{}) error {
	s.rw.RLock()
	defer s.rw.RUnlock()

	if v, ok := s.items[key]; ok {
		if !v.isExpired() {
			return cache.Decode(v.value, dest)
		} else {
			delete(s.items, key)
		}
	}

	return cache.ErrNotFound
}

func (s *Store) Put(key string, value interface{}, ttl time.Duration) error {
	s.rw.Lock()
	defer s.rw.Unlock()

	s.items[key] = &item{
		value:   value,
		expired: time.Now().Add(ttl),
	}

	return nil
}

func (s *Store) Increment(key string, value int) (int, error) {
	s.rw.Lock()
	defer s.rw.Unlock()

	if v, ok := s.items[key]; ok {
		if !v.isExpired() {
			if i, ok := v.value.(int); ok {
				v.value = i + value
				return v.value.(int), nil
			} else {
				return 0, cache.ErrInvalidDataType
			}
		}
	}

	s.items[key] = &item{
		value: value,
	}

	return value, nil
}

func (s *Store) Decrement(key string, value int) (int, error) {
	return s.Increment(key, -value)
}

func (s *Store) Forever(key string, value interface{}) bool {
	s.rw.Lock()
	defer s.rw.Unlock()

	s.items[key] = &item{
		value: value,
	}

	return true
}

func (s *Store) Forget(key string) error {
	s.rw.Lock()
	defer s.rw.Unlock()

	delete(s.items, key)

	return nil
}

func (s *Store) Flush() error {
	s.rw.Lock()
	defer s.rw.Unlock()

	s.items = make(map[string]*item)

	return nil
}

func (s *Store) GetPrefix() string {
	return ""
}
