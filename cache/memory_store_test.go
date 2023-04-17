package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMemoryStore(t *testing.T) {
	m := NewMemoryStore()

	assert.True(t, m.Put("foo", "bar", time.Now().Add(time.Second*10)))
	assert.Equal(t, "bar", m.Get("foo"))
}
