package cache

import (
	"github.com/go-fires/fires/x/contracts/cache"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func createMemoryStore() *MemoryStore {
	return NewMemoryStore(&MemoryStoreConfig{
		LruTick: time.Minute * 1,
	})
}

func createRedisStore() *RedisStore {
	r := NewRedisStore(redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	}))

	r.Flush()

	return r
}

func TestStore_Putget(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"memory", createMemoryStore()},
		{"redis", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, tt.store.Put("foo", "bar", time.Second*1))
			var foo string
			assert.Nil(t, tt.store.Get("foo", &foo))
			assert.Equal(t, "bar", foo)

			time.Sleep(time.Second * 2)

			var foo2 string
			assert.Error(t, tt.store.Get("foo", &foo2))
			assert.Equal(t, "", foo2)

			type User struct {
				Name string
				Age  int
			}

			var user User
			assert.True(t, tt.store.Put("user", User{"foo", 18}, time.Second*1))
			assert.Nil(t, tt.store.Get("user", &user))
			assert.Equal(t, "foo", user.Name)
			assert.Equal(t, 18, user.Age)
		})
	}

}

func TestStore_IncrAndDecr(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"memory", createMemoryStore()},
		{"redis", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, 1, tt.store.Increment("foo", 1))
			assert.Equal(t, 6, tt.store.Increment("foo", 5))
			assert.Equal(t, 14, tt.store.Increment("foo", 8))

			assert.Equal(t, 13, tt.store.Decrement("foo", 1))
			assert.Equal(t, 8, tt.store.Decrement("foo", 5))
			assert.Equal(t, 0, tt.store.Decrement("foo", 8))

			// test overflow
			assert.True(t, tt.store.Put("foo", "bar", time.Second*10))
			var foo string
			assert.Nil(t, tt.store.Get("foo", &foo))
			assert.Equal(t, "bar", foo)
			assert.Equal(t, 0, tt.store.Increment("foo", 1))
		})
	}

}

func TestStore_Forever(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"memory", createMemoryStore()},
		{"redis", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, tt.store.Forever("foo", "bar"))
			var foo string
			assert.Nil(t, tt.store.Get("foo", &foo))
			assert.Equal(t, "bar", foo)
		})
	}
}

func TestStore_Forget(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"memory", createMemoryStore()},
		{"redis", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var foo string
			tt.store.Put("foo", "bar", time.Second*1)
			assert.Nil(t, tt.store.Get("foo", &foo))
			assert.Equal(t, "bar", foo)
			assert.True(t, tt.store.Forget("foo"))
			assert.False(t, tt.store.Has("foo"))

			assert.False(t, tt.store.Forget("foo2"))
		})
	}
}

func TestStore_Flush(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"memory", createMemoryStore()},
		{"redis", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var foo, foo2 string
			tt.store.Put("foo", "bar", time.Second*1)
			tt.store.Put("foo2", "bar2", time.Second*1)
			assert.Nil(t, tt.store.Get("foo", &foo))
			assert.Nil(t, tt.store.Get("foo2", &foo2))

			assert.True(t, tt.store.Flush())
			var foo3, foo4 string
			assert.Error(t, tt.store.Get("foo", &foo3))
			assert.Error(t, tt.store.Get("foo2", &foo4))
		})
	}
}
