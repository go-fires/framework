package cache

import (
	"errors"
	"time"
)

var (
	ErrNotFound        = errors.New("cache: key not found")
	ErrInvalidDataType = errors.New("cache: invalid data type")
)

type Store interface {
	Has(key string) bool

	Get(key string, dest interface{}) error

	Put(key string, value interface{}, ttl time.Duration) error

	Increment(key string, value int) (int, error)

	Decrement(key string, value int) (int, error)

	Forever(key string, value interface{}) bool

	Forget(key string) error

	Flush() error

	GetPrefix() string
}

type StoreAddable interface {
	Add(key string, value interface{}, ttl time.Duration) (bool, error)
}

type StorePullable interface {
	Pull(key string, dest interface{}) error
}
