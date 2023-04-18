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
