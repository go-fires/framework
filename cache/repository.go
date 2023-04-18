package cache

import (
	"github.com/go-fires/framework/contracts/cache"
	"github.com/go-fires/framework/support"
	"time"
)

type Respository struct {
	cache.Store
}

var _ cache.Respository = (*Respository)(nil)

func NewRespository(store cache.Store) *Respository {
	return &Respository{
		Store: store,
	}
}

func (r *Respository) Has(key string) bool {
	return r.Get(key) != nil
}

func (r *Respository) Missing(key string) bool {
	return !r.Has(key)
}

func (r *Respository) Pull(key string) interface{} {
	return support.Tap(r.Get(key), func(value interface{}) {
		r.Forget(key)
	})
}

func (r *Respository) Set(key string, value interface{}, ttl time.Duration) bool {
	return r.Put(key, value, ttl)
}

func (r *Respository) Add(key string, value interface{}, ttl time.Duration) bool {
	// if the store supports the add method, we'll just call that and return
	if s, ok := r.Store.(cache.StoreAddable); ok {
		return s.Add(key, value, ttl)
	}

	// otherwise we'll just simulate it by checking for an existing value and
	if r.Has(key) {
		return false
	}

	return r.Set(key, value, ttl)
}

func (r *Respository) Remember(key string, ttl time.Duration, callback func() interface{}) interface{} {
	value := r.Get(key)
	if value != nil {
		return value
	}

	value = callback()

	r.Set(key, value, ttl)

	return value
}

func (r *Respository) RememberForever(key string, callback func() interface{}) interface{} {
	return r.Remember(key, 0, callback)
}

func (r *Respository) Delete(key string) bool {
	return r.Forget(key)
}

func (r *Respository) Clear() bool {
	return r.Flush()
}
