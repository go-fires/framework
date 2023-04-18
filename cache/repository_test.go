package cache

import (
	"github.com/go-fires/framework/contracts/cache"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRespository_Base(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"memory", createMemoryStore()},
		{"redis", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRespository(tt.store)

			assert.True(t, r.Put("foo", "bar", time.Second*1))
			assert.True(t, r.Has("foo"))
			assert.False(t, r.Missing("foo"))

			time.Sleep(time.Second * 2)
			assert.False(t, r.Has("foo"))
			assert.True(t, r.Missing("foo"))
		})
	}
}

func TestRespository_Pull(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"memory", createMemoryStore()},
		{"redis", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRespository(tt.store)

			var foo string
			r.Put("foo", "bar", time.Second*100)
			assert.Nil(t, r.Pull("foo", &foo))
			assert.Equal(t, "bar", foo)
			assert.False(t, r.Has("foo"))
		})
	}
}

func TestRespository_Set(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"memory", createMemoryStore()},
		{"redis", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRespository(tt.store)

			var foo string
			assert.True(t, r.Set("foo", "bar", time.Second*1))
			assert.Nil(t, r.Get("foo", &foo))
			assert.Equal(t, "bar", foo)
		})
	}
}

func TestRespository_Add(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"redis", createMemoryStore()},
		{"memory", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRespository(tt.store)

			var foo string
			assert.True(t, r.Add("foo", "bar", time.Second*1))
			assert.Nil(t, r.Get("foo", &foo))
			assert.Equal(t, "bar", foo)
			assert.False(t, r.Add("foo", "bar", time.Second*1))
		})
	}
}

func TestRespository_Remember(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"redis", createMemoryStore()},
		{"memory", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRespository(tt.store)

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

func TestRespository_RememberForever(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"redis", createMemoryStore()},
		{"memory", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRespository(tt.store)

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

func TestRespository_Delete(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"redis", createMemoryStore()},
		{"memory", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRespository(tt.store)

			r.Put("foo", "bar", time.Second*1)
			assert.True(t, r.Has("foo"))
			assert.True(t, r.Delete("foo"))
			assert.False(t, r.Has("foo"))
			assert.False(t, r.Delete("foo"))
		})
	}
}

func TestRespository_Clear(t *testing.T) {
	for _, tt := range []struct {
		name  string
		store cache.Store
	}{
		{"redis", createMemoryStore()},
		{"memory", createRedisStore()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRespository(tt.store)

			r.Put("foo", "bar", time.Second*1)
			r.Put("baz", "bar", time.Second*1)
			assert.True(t, r.Has("foo"))
			assert.True(t, r.Has("baz"))
			assert.True(t, r.Clear())
			assert.False(t, r.Has("foo"))
			assert.False(t, r.Has("baz"))
		})
	}
}
