package cache

import (
	"time"
)

type Respository interface {
	Store

	Has(key string) bool
	Missing(key string) bool
	Pull(key string) interface{}
	Set(key string, value interface{}, ttl time.Duration) bool
	Add(key string, value interface{}, ttl time.Duration) bool
	Remember(key string, ttl time.Duration, callback func() interface{}) interface{}
	RememberForever(key string, callback func() interface{}) interface{}
	Delete(key string) bool
	Clear() bool
}
