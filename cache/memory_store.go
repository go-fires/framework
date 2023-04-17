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

func (r *record) isExpired() bool {
	return !r.expired.IsZero() && !r.expired.After(time.Now())
}

type MemoryStore struct {
	records map[string]*record

	rw sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	m := &MemoryStore{
		records: make(map[string]*record, 1024),
	}

	go m.lrucache()

	return m
}

var _ cache.Store = (*MemoryStore)(nil)

func (m *MemoryStore) Get(key string) interface{} {
	m.rw.RLock()
	defer m.rw.RUnlock()

	if v, ok := m.records[key]; ok {
		if !v.isExpired() {
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

	m.records[key] = &record{
		value:   value,
		expired: expired,
	}

	return true
}

func (m *MemoryStore) Increment(key string, value int) int {
	m.rw.RLock()
	defer m.rw.RUnlock()

	if v, ok := m.records[key]; ok {
		if !v.isExpired() {
			if _, ok := v.value.(int); ok {
				v.value = v.value.(int) + value
			} else {
				v.value = value
			}

			return v.value.(int)
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
	m.rw.Lock()
	defer m.rw.Unlock()

	if _, ok := m.records[key]; ok {
		delete(m.records, key)

		return true
	}

	return false
}

func (m *MemoryStore) Flush() bool {
	m.rw.Lock()
	defer m.rw.Unlock()

	m.records = make(map[string]*record, 1024)

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

func (m *MemoryStore) lrucache() {
	for {
		select {
		case <-time.Tick(time.Minute):
			for k, v := range m.getRecords() {
				if v.isExpired() {
					m.Forget(k)
				}
			}
		}
	}
}

func (m *MemoryStore) getRecords() map[string]*record {
	return m.records
}
