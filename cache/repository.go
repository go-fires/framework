package cache

import (
	"time"
)

type Repository struct {
	Store
}

func NewRepository(store Store) *Repository {
	return &Repository{
		Store: store,
	}
}

func (r *Repository) Missing(key string) bool {
	return !r.Has(key)
}

func (r *Repository) Pull(key string, dest interface{}) error {
	if pullable, ok := r.Store.(StorePullable); ok {
		return pullable.Pull(key, dest)
	}

	if err := r.Get(key, dest); err != nil {
		return err
	}

	return r.Forget(key)
}

func (r *Repository) Set(key string, value interface{}, ttl time.Duration) error {
	return r.Put(key, value, ttl)
}

func (r *Repository) Add(key string, value interface{}, ttl time.Duration) (bool, error) {
	// instance of StoreAddable
	if addable, ok := r.Store.(StoreAddable); ok {
		return addable.Add(key, value, ttl)
	}

	if r.Has(key) {
		return false, nil
	}

	if err := r.Put(key, value, ttl); err != nil {
		return false, err
	}

	return true, nil
}

func (r *Repository) Remember(key string, dest interface{}, ttl time.Duration, callback func() interface{}) error {
	if r.Has(key) {
		return r.Get(key, dest)
	}

	value := callback()
	if err := r.Put(key, value, ttl); err != nil {
		return err
	}

	return Decode(value, dest)
}

func (r *Repository) RememberForever(key string, dest interface{}, callback func() interface{}) error {
	if r.Has(key) {
		return r.Get(key, dest)
	}

	value := callback()
	if err := r.Put(key, value, 0); err != nil {
		return err
	}

	return Decode(value, dest)
}

func (r *Repository) Delete(key string) error {
	return r.Forget(key)
}

func (r *Repository) Clear() error {
	return r.Flush()
}
