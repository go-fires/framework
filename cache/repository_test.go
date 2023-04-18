package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRespository_Base(t *testing.T) {
	r := NewRespository(NewMemoryStore())

	assert.True(t, r.Put("foo", "bar", time.Second*1))
	assert.True(t, r.Has("foo"))
	assert.False(t, r.Missing("foo"))

	time.Sleep(time.Second * 2)
	assert.False(t, r.Has("foo"))
	assert.True(t, r.Missing("foo"))
}

func TestRespository_Pull(t *testing.T) {
	r := NewRespository(NewMemoryStore())

	r.Put("foo", "bar", time.Second*1)
	assert.Equal(t, "bar", r.Pull("foo"))
	assert.Nil(t, r.Get("foo"))
}

func TestRespository_Set(t *testing.T) {
	r := NewRespository(NewMemoryStore())

	assert.True(t, r.Set("foo", "bar", time.Second*1))
	assert.Equal(t, "bar", r.Get("foo"))
}

func TestRespository_Add(t *testing.T) {
	r := NewRespository(NewMemoryStore())

	assert.True(t, r.Add("foo", "bar", time.Second*1))
	assert.Equal(t, "bar", r.Get("foo"))
	assert.False(t, r.Add("foo", "bar", time.Second*1))
}

func TestRespository_Remember(t *testing.T) {
	r := NewRespository(NewMemoryStore())

	assert.Equal(t, "bar", r.Remember("foo", time.Second*1, func() interface{} {
		return "bar"
	}))
	assert.Equal(t, "bar", r.Get("foo"))

	assert.Equal(t, "bar", r.Remember("foo", time.Second*1, func() interface{} {
		return "baz"
	}))
	assert.Equal(t, "bar", r.Get("foo"))
}

func TestRespository_RememberForever(t *testing.T) {
	r := NewRespository(NewMemoryStore())

	assert.Equal(t, "bar", r.RememberForever("foo", func() interface{} {
		return "bar"
	}))
	assert.Equal(t, "bar", r.Get("foo"))

	assert.Equal(t, "bar", r.RememberForever("foo", func() interface{} {
		return "baz"
	}))
	assert.Equal(t, "bar", r.Get("foo"))
}

func TestRespository_Delete(t *testing.T) {
	r := NewRespository(NewMemoryStore())

	r.Put("foo", "bar", time.Second*1)
	assert.True(t, r.Has("foo"))
	assert.True(t, r.Delete("foo"))
	assert.False(t, r.Has("foo"))
	assert.False(t, r.Delete("foo"))
}

func TestRespository_Clear(t *testing.T) {
	r := NewRespository(NewMemoryStore())

	r.Put("foo", "bar", time.Second*1)
	r.Put("baz", "bar", time.Second*1)
	assert.True(t, r.Has("foo"))
	assert.True(t, r.Has("baz"))
	assert.True(t, r.Clear())
	assert.False(t, r.Has("foo"))
	assert.False(t, r.Has("baz"))
}
