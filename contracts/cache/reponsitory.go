package cache

import (
	"time"
)

type Repository interface {
	Store

	Missing(key string) bool
	Pull(key string, value interface{}) error
	Set(key string, value interface{}, ttl time.Duration) bool
	Add(key string, value interface{}, ttl time.Duration) bool
	Remember(key string, value interface{}, ttl time.Duration, callback func() interface{}) error
	RememberForever(key string, value interface{}, callback func() interface{}) error
	Delete(key string) bool
	Clear() bool
}
