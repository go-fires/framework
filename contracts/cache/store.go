package cache

import "time"

type Store interface {
	Get(key string) interface{}

	Put(key string, value interface{}, ttl time.Duration) bool

	Increment(key string, value int) int

	Decrement(key string, value int) int

	Forever(key string, value interface{}) bool

	Forget(key string) bool

	Flush() bool

	GetPrefix() string
}

type StoreAddable interface {
	Add(key string, value interface{}, ttl time.Duration) bool
}
