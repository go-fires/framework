package cache

import (
	"github.com/go-fires/framework/contracts/cache"
	"github.com/go-fires/framework/support/helper"
	"sync"
	"time"
)

type Repository struct {
	cache.Store

	mu sync.Mutex
}

var _ cache.Repository = (*Repository)(nil)

func NewRepository(store cache.Store) *Repository {
	return &Repository{
		Store: store,
	}
}

func (r *Repository) Missing(key string) bool {
	return !r.Has(key)
}

func (r *Repository) Pull(key string, value interface{}) error {
	if s, ok := r.Store.(cache.StorePullable); ok {
		return s.Pull(key, value)
	}

	switch helper.Tap(r.Get(key, value), func(value interface{}) {
		r.Forget(key)
	}).(type) {
	case nil:
		return nil
	case error:
		return value.(error)
	default:
		return cache.ErrUnknown
	}
}

func (r *Repository) Set(key string, value interface{}, ttl time.Duration) bool {
	return r.Put(key, value, ttl)
}

func (r *Repository) Add(key string, value interface{}, ttl time.Duration) bool {
	// if the store supports the add method, we'll just call that and return
	if s, ok := r.Store.(cache.StoreAddable); ok {
		return s.Add(key, value, ttl)
	}

	// otherwise we'll just simulate it by checking for an existing value and
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.Has(key) {
		return false
	}

	return r.Set(key, value, ttl)
}

// Remember calls the given callback and stores the result if the key does not exist.
// Example:
//
//	var value string
//	cache.Remember("key", &value, time.Minute, func() interface{} {
//		return "value"
//	})
func (r *Repository) Remember(key string, value interface{}, ttl time.Duration, callback func() interface{}) error {
	if nil == r.Get(key, value) {
		return nil
	}

	return helper.ValueOf(helper.Tap(callback(), func(value interface{}) {
		r.Set(key, value, ttl)
	}), value)
}

func (r *Repository) RememberForever(key string, value interface{}, callback func() interface{}) error {
	return r.Remember(key, value, 0, callback)
}

func (r *Repository) Delete(key string) bool {
	return r.Forget(key)
}

func (r *Repository) Clear() bool {
	return r.Flush()
}
