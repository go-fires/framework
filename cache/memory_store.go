package cache

import (
	"github.com/go-fires/framework/contracts/cache"
	"sync"
	"time"
)

type record struct {
	value   interface{}
	expired time.Time
}

type MemoryStore struct {
	records map[string]record

	rw sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		records: make(map[string]record, 1024),
	}
}

var _ cache.Store = (*MemoryStore)(nil)

func (m *MemoryStore) Get(key string) interface{} {
	if v, ok := m.records[key]; ok {
		if v.expired.IsZero() || v.expired.After(time.Now()) {
			return v.value
		} else {
			delete(m.records, key)
		}
	}

	return nil
}

func (m *MemoryStore) Put(key string, value interface{}, expired time.Time) bool {
	m.rw.Lock()
	defer m.rw.Unlock()

	m.records[key] = record{
		value:   value,
		expired: expired,
	}

	return true
}

func (m *MemoryStore) Increment(key string, value int) int {
	if v, ok := m.records[key]; ok {
		if v.expired.IsZero() || v.expired.After(time.Now()) {
			if i, ok := v.value.(int); ok {
				i += value
				m.Put(key, i, v.expired)

				return i
			}
		} else {
			delete(m.records, key)
		}
	}

	m.Put(key, value, time.Time{}) // forever

	return value
}

func (m *MemoryStore) Decrement(key string, value int) int {
	return m.Increment(key, -value)
}

func (m *MemoryStore) Forever(key string, value interface{}) bool {
	return m.Put(key, value, time.Time{})
}

func (m *MemoryStore) Forget(key string) bool {
	if _, ok := m.records[key]; ok {
		delete(m.records, key)

		return true
	}

	return false
}

func (m *MemoryStore) Flush() bool {
	m.rw.Lock()
	defer m.rw.Unlock()

	m.records = make(map[string]record, 1024)

	return true
}

func (m *MemoryStore) Has(key string) bool {
	if _, ok := m.records[key]; ok {
		return true
	}

	return false
}

func (m *MemoryStore) GetPrefix() string {
	return ""
}
