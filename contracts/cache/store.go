package cache

import "time"

type Store interface {
	Get(key string) interface{}

	Put(key string, value interface{}, expired time.Time) bool

	Increment(key string, value int) int

	Decrement(key string, value int) int

	Forever(key string, value interface{}) bool

	Forget(key string) bool

	Flush() bool

	Has(key string) bool

	GetPrefix() string
}
