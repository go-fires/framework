package cache

import (
	"context"
	"github.com/go-fires/framework/contracts/cache"
	"github.com/go-fires/framework/contracts/support"
	"github.com/go-fires/framework/support/serializer"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

type RedisStoreConfig struct {
	Connection string // see redis/config.go (use Connections key)
	Prefix     string
	Serializer support.Serializable
}

var _ cache.StoreConfigable = (*RedisStoreConfig)(nil)

func (r *RedisStoreConfig) GetConnection() string {
	if r.Connection == "" {
		return "default"
	}

	return r.Connection
}

func (r *RedisStoreConfig) GetPrefix() string {
	if r.Prefix == "" {
		return "cache:"
	}

	return r.Prefix
}

func (r *RedisStoreConfig) GetSerializer() support.Serializable {
	if r.Serializer == nil {
		return &serializer.JsonSerializer{}
	}

	return r.Serializer
}

type RedisStore struct {
	redis        redis.Cmdable
	serializable support.Serializable

	prefix string
}

var _ cache.Store = (*RedisStore)(nil)
var _ cache.StoreAddable = (*RedisStore)(nil)  // support for add method
var _ cache.StorePullable = (*RedisStore)(nil) // support for pull method

type RedisStoreOption func(*RedisStore)

func NewRedisStore(redis redis.Cmdable, opts ...RedisStoreOption) *RedisStore {
	r := &RedisStore{
		redis:        redis,
		serializable: &serializer.JsonSerializer{}, // default serializer
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func WithRedisStorePrefix(prefix string) RedisStoreOption {
	return func(r *RedisStore) {
		r.SetPrefix(prefix)
	}
}

func WithRedisStoreSerializable(serializable support.Serializable) RedisStoreOption {
	return func(r *RedisStore) {
		r.serializable = serializable
	}
}

func (r *RedisStore) Has(key string) bool {
	if result := r.redis.Exists(ctx, r.prefix+key); result.Err() != nil {
		return false
	} else {
		return result.Val() > 0
	}
}

func (r *RedisStore) Get(key string, value interface{}) error {
	if result := r.redis.Get(ctx, r.prefix+key); result.Err() != nil {
		return result.Err()
	} else {
		return r.unserialize(result.Val(), value)
	}
}

func (r *RedisStore) Put(key string, value interface{}, ttl time.Duration) bool {
	serialized, err := r.serialize(value)
	if err != nil {
		return false
	}

	if result := r.redis.Set(ctx, r.prefix+key, serialized, ttl); result.Err() != nil {
		return false
	} else {
		return true
	}
}

func (r *RedisStore) Increment(key string, value int) int {
	if result := r.redis.IncrBy(ctx, r.prefix+key, int64(value)); result.Err() != nil {
		return 0
	} else {
		return int(result.Val())
	}
}

func (r *RedisStore) Decrement(key string, value int) int {
	if result := r.redis.DecrBy(ctx, r.prefix+key, int64(value)); result.Err() != nil {
		return 0
	} else {
		return int(result.Val())
	}
}

func (r *RedisStore) Pull(key string, value interface{}) error {
	if result := r.redis.GetDel(ctx, r.prefix+key); result.Err() != nil {
		return result.Err()
	} else {
		return r.unserialize(result.Val(), value)
	}
}

func (r *RedisStore) Forever(key string, value interface{}) bool {
	return r.Put(r.prefix+key, value, 0)
}

func (r *RedisStore) Forget(key string) bool {
	if result := r.redis.Del(ctx, r.prefix+key); result.Err() != nil {
		return false
	} else {
		return result.Val() > 0
	}
}

func (r *RedisStore) Add(key string, value interface{}, ttl time.Duration) bool {
	serialized, err := r.serialize(value)
	if err != nil {
		return false
	}

	if result := r.redis.SetNX(ctx, r.prefix+key, serialized, ttl); result.Err() != nil {
		return false
	} else {
		return result.Val()
	}
}

func (r *RedisStore) Flush() bool {
	if result := r.redis.FlushDB(ctx); result.Err() != nil {
		return false
	} else {
		return true
	}
}

func (r *RedisStore) SetPrefix(prefix string) {
	if prefix != "" {
		prefix = prefix + ":"
	}

	r.prefix = prefix
}

func (r *RedisStore) GetPrefix() string {
	return r.prefix
}

func (r *RedisStore) serialize(value interface{}) (string, error) {
	if result, err := r.serializable.Serialize(value); err != nil {
		return "", err
	} else {
		return string(result), nil
	}
}

func (r *RedisStore) unserialize(value string, dest interface{}) error {
	return r.serializable.Unserialize([]byte(value), dest)
}
