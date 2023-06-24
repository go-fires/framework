package cache

import (
	"sync"
	"time"

	"github.com/go-fires/fires/contracts/cache"
	"github.com/go-fires/fires/support/helper"
)

type record struct {
	value   interface{}
	expired time.Time
}

func (r *record) isExpired() bool {
	return !r.expired.IsZero() && !r.expired.After(time.Now())
}

type MemoryStoreConfig struct {
	LruSize int // todo: lru size
	LruTick time.Duration
}

var _ StoreConfigable = (*RedisStoreConfig)(nil)

func (c *MemoryStoreConfig) GetLruSize() int {
	if c.LruSize == 0 {
		return 1000
	}

	return c.LruSize
}

func (c *MemoryStoreConfig) GetLruTick() time.Duration {
	if c.LruTick == 0 {
		return time.Minute
	}

	return c.LruTick
}

type MemoryStore struct {
	config  *MemoryStoreConfig
	records *sync.Map

	mu sync.Mutex
}

func NewMemoryStore(config *MemoryStoreConfig) *MemoryStore {
	m := &MemoryStore{
		config:  config,
		records: &sync.Map{},
	}

	go m.lrucache()

	return m
}

var _ cache.Store = (*MemoryStore)(nil)

func (m *MemoryStore) Has(key string) bool {
	if v, ok := m.records.Load(key); ok {
		if !v.(*record).isExpired() {
			return true
		} else {
			m.records.Delete(key)
		}
	}

	return false
}

// Get gets a value from the cache.
func (m *MemoryStore) Get(key string, value interface{}) error {
	if v, ok := m.records.Load(key); ok {
		if !v.(*record).isExpired() {
			return helper.ValueOf(v.(*record).value, value)
		} else {
			m.records.Delete(key)
		}
	}

	return cache.ErrKeyNotFound
}

// Put puts a value into the cache for a given number of minutes.
func (m *MemoryStore) Put(key string, value interface{}, ttl time.Duration) bool {
	m.records.Store(key, &record{
		value:   value,
		expired: m.getExpired(ttl),
	})

	return true
}

// Increment increments the value of an item in the cache.
func (m *MemoryStore) Increment(key string, value int) int {
	m.mu.Lock()
	defer m.mu.Unlock()

	if v, ok := m.records.Load(key); ok {
		if !v.(*record).isExpired() {
			if _, ok := v.(*record).value.(int); ok {
				v.(*record).value = v.(*record).value.(int) + value

				return v.(*record).value.(int)
			} else {
				return 0 // not int
			}
		} else {
			m.records.Delete(key)
		}
	}

	m.Forever(key, value) // forever

	return value
}

// Decrement decrements the value of an item in the cache.
func (m *MemoryStore) Decrement(key string, value int) int {
	return m.Increment(key, -value)
}

// Forever stores an item in the cache indefinitely.
func (m *MemoryStore) Forever(key string, value interface{}) bool {
	return m.Put(key, value, 0)
}

// Forget removes an item from the cache.
func (m *MemoryStore) Forget(key string) bool {
	if _, ok := m.records.Load(key); ok {
		m.records.Delete(key)
		return true
	}

	return false
}

// Flush clears the cache.
func (m *MemoryStore) Flush() bool {
	m.records = &sync.Map{}

	return true
}

// GetPrefix returns the prefix for the cache.
func (m *MemoryStore) GetPrefix() string {
	return ""
}

// lrucache is a goroutine that periodically removes expired items from the cache.
func (m *MemoryStore) lrucache() {
	for range time.Tick(m.config.GetLruTick()) {
		m.records.Range(func(key, value interface{}) bool {
			if value.(*record).isExpired() {
				m.Forget(key.(string))
			}

			return true
		})
	}
}

// getExpired returns the time when the item should expire.
func (m *MemoryStore) getExpired(ttl time.Duration) time.Time {
	if ttl == 0 {
		return time.Time{}
	}

	return time.Now().Add(ttl)
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
