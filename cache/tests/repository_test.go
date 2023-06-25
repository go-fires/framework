package tests

import (
	"github.com/go-fires/fires/cache"
	"github.com/go-fires/fires/cache/store/memory"
	"github.com/go-fires/fires/cache/store/redis"
	rds "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func createMemoryStore() *memory.Store {
	return memory.New()
}

func createRedisStore() *redis.Store {
	return redis.New(rds.NewClient(&rds.Options{
		Addr: "localhost:6379",
	}))
}

func TestRepository_Base(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"memory", createMemoryStore()},
		{"redis", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := cache.NewRepository(tt.store)

			assert.NoError(t, r.Put("foo", "bar", time.Second*1))
			assert.True(t, r.Has("foo"))
			assert.False(t, r.Missing("foo"))

			time.Sleep(time.Second * 2)
			assert.False(t, r.Has("foo"))
			assert.True(t, r.Missing("foo"))
		})
	}
}

func TestRepository_Pull(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"memory", createMemoryStore()},
		{"redis", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := cache.NewRepository(tt.store)

			var foo string
			r.Put("foo", "bar", time.Second*100)
			assert.Nil(t, r.Pull("foo", &foo))
			assert.Equal(t, "bar", foo)
			assert.False(t, r.Has("foo"))
		})
	}
}

func TestRepository_Set(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"memory", createMemoryStore()},
		{"redis", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := cache.NewRepository(tt.store)

			var foo string
			assert.NoError(t, r.Set("foo", "bar", time.Second*1))
			assert.Nil(t, r.Get("foo", &foo))
			assert.Equal(t, "bar", foo)
		})
	}
}

func TestRepository_Add(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"memory", createMemoryStore()},
		{"redis", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := cache.NewRepository(tt.store)

			var foo string
			b1, err1 := r.Add("foo", "bar", time.Second*1)
			assert.NoError(t, err1)
			assert.True(t, b1)
			assert.Nil(t, r.Get("foo", &foo))
			assert.Equal(t, "bar", foo)

			b2, err2 := r.Add("foo", "baz", time.Second*1)
			assert.NoError(t, err2)
			assert.False(t, b2)
		})
	}
}

func TestRepository_Remember(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"memory", createMemoryStore()},
		{"redis", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := cache.NewRepository(tt.store)

			var bar string
			assert.Nil(t, r.Remember("foo", &bar, time.Second*1, func() interface{} {
				return "bar"
			}))
			assert.Equal(t, "bar", bar)

			assert.Nil(t, r.Remember("foo", &bar, time.Second*1, func() interface{} {
				return "baz"
			}))
			assert.Equal(t, "bar", bar)
		})
	}
}

func TestRepository_RememberForever(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"memory", createMemoryStore()},
		// {"memory", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := cache.NewRepository(tt.store)

			var foo string
			assert.Nil(t, r.RememberForever("foo", &foo, func() interface{} {
				return "bar"
			}))
			assert.Equal(t, "bar", foo)

			assert.Nil(t, r.RememberForever("foo", &foo, func() interface{} {
				return "baz"
			}))
			assert.Equal(t, "bar", foo)
		})
	}
}

func TestRepository_Delete(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"memory", createMemoryStore()},
		{"redis", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := cache.NewRepository(tt.store)

			r.Put("foo", "bar", time.Second*1)
			assert.True(t, r.Has("foo"))
			assert.NoError(t, r.Delete("foo"))
			assert.False(t, r.Has("foo"))
		})
	}
}

func TestRepository_Clear(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"redis", createMemoryStore()},
		{"redis", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := cache.NewRepository(tt.store)

			r.Put("foo", "bar", time.Second*1)
			r.Put("baz", "bar", time.Second*1)
			assert.True(t, r.Has("foo"))
			assert.True(t, r.Has("baz"))
			assert.NoError(t, r.Clear())
			assert.False(t, r.Has("foo"))
			assert.False(t, r.Has("baz"))
		})
	}
}
