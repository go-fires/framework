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
	records *sync.Map

	mu sync.Mutex
}

func NewMemoryStore() *MemoryStore {
	m := &MemoryStore{
		records: &sync.Map{},
	}

	go m.lrucache()

	return m
}

var _ cache.Store = (*MemoryStore)(nil)

func (m *MemoryStore) Get(key string) interface{} {
	if v, ok := m.records.Load(key); ok {
		if !v.(*record).isExpired() {
			return v.(*record).value
		} else {
			m.records.Delete(key)
		}
	}

	return nil
}

func (m *MemoryStore) Put(key string, value interface{}, expired time.Time) bool {
	m.records.Store(key, &record{
		value:   value,
		expired: expired,
	})

	return true
}

func (m *MemoryStore) Increment(key string, value int) int {
	m.mu.Lock()
	defer m.mu.Unlock()

	if v, ok := m.records.Load(key); ok {
		if !v.(*record).isExpired() {
			if _, ok := v.(*record).value.(int); ok {
				v.(*record).value = v.(*record).value.(int) + value
			} else {
				v.(*record).value = value
			}

			return v.(*record).value.(int)
		} else {
			m.records.Delete(key)
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
	if _, ok := m.records.Load(key); ok {
		m.records.Delete(key)
		return true
	}

	return false
}

func (m *MemoryStore) Flush() bool {
	m.records = &sync.Map{}

	return true
}

func (m *MemoryStore) Has(key string) bool {
	if _, ok := m.records.Load(key); ok {
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
			m.records.Range(func(key, value interface{}) bool {
				if value.(*record).isExpired() {
					m.Forget(key.(string))
				}

				return true
			})
		}
	}
}

// Length returns the number of items in the cache.
// This is slow, so it should only be used for debugging.
func (m *MemoryStore) Length() int {
	var length int

	m.records.Range(func(key, value interface{}) bool {
		length++

		return true
	})

	return length
}
