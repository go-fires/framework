package redis

import (
	"context"
	"github.com/go-fires/fires/cache"
	"github.com/go-fires/fires/serializer"
	"github.com/redis/go-redis/v9"
	"time"
)

type Option func(*Store)

type Store struct {
	redis redis.Cmdable

	prefix     string
	context    context.Context
	serializer serializer.Serializer
}

var _ cache.Store = (*Store)(nil)
var _ cache.StoreAddable = (*Store)(nil)  // Support for add method
var _ cache.StorePullable = (*Store)(nil) // Support for pull method

func New(redis redis.Cmdable, opts ...Option) *Store {
	s := &Store{
		redis:      redis,
		serializer: serializer.Json,
		context:    context.Background(),
	}

	return s.With(opts...)
}

func WithPrefix(prefix string) Option {
	return func(s *Store) {
		if prefix == "" {
			return
		}

		s.prefix = prefix + ":"
	}
}

func WithSerializer(serializer serializer.Serializer) Option {
	return func(s *Store) {
		if serializer == nil {
			return
		}

		s.serializer = serializer
	}
}

func WithContext(ctx context.Context) Option {
	return func(s *Store) {
		if ctx == nil {
			return
		}

		s.context = ctx
	}
}

func (s *Store) With(opts ...Option) *Store {
	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Store) Has(key string) bool {
	if result := s.redis.Exists(s.context, s.prefix+key); result.Err() != nil {
		return false
	} else {
		return result.Val() > 0
	}
}

func (s *Store) Get(key string, dest interface{}) error {
	if result := s.redis.Get(s.context, s.prefix+key); result.Err() != nil {
		return result.Err()
	} else {
		return s.serializer.Unserialize([]byte(result.Val()), dest)
	}
}

func (s *Store) Put(key string, value interface{}, ttl time.Duration) error {
	if data, err := s.serializer.Serialize(value); err != nil {
		return err
	} else {
		return s.redis.Set(s.context, s.prefix+key, data, ttl).Err()
	}
}

func (s *Store) Increment(key string, value int) (int, error) {
	if result := s.redis.IncrBy(s.context, s.prefix+key, int64(value)); result.Err() != nil {
		return 0, result.Err()
	} else {
		return int(result.Val()), nil
	}
}

func (s *Store) Decrement(key string, value int) (int, error) {
	if result := s.redis.DecrBy(s.context, s.prefix+key, int64(value)); result.Err() != nil {
		return 0, result.Err()
	} else {
		return int(result.Val()), nil
	}
}

func (s *Store) Forever(key string, value interface{}) bool {
	return s.redis.Set(s.context, s.prefix+key, value, 0).Err() == nil
}

func (s *Store) Forget(key string) error {
	return s.redis.Del(s.context, s.prefix+key).Err()
}

func (s *Store) Flush() error {
	return s.redis.FlushDB(s.context).Err()
}

func (s *Store) GetPrefix() string {
	return s.prefix
}

func (s *Store) Add(key string, value interface{}, ttl time.Duration) (bool, error) {
	if data, err := s.serializer.Serialize(value); err != nil {
		return false, err
	} else {
		return s.redis.SetNX(s.context, s.prefix+key, data, ttl).Result()
	}
}

func (s *Store) Pull(key string, dest interface{}) error {
	if result := s.redis.Get(s.context, s.prefix+key); result.Err() != nil {
		return result.Err()
	} else {
		if err := s.serializer.Unserialize([]byte(result.Val()), dest); err != nil {
			return err
		} else {
			return s.Forget(key)
		}
	}
}
