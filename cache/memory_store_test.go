package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMemoryStore_PutGet(t *testing.T) {
	m := NewMemoryStore()

	assert.True(t, m.Put("foo", "bar", time.Now().Add(time.Second*1)))
	assert.Equal(t, "bar", m.Get("foo").(string))

	time.Sleep(time.Second * 2)
	assert.Nil(t, m.Get("foo"))
}

func TestMemoryStore_IncrAndDecr(t *testing.T) {
	m := NewMemoryStore()

	assert.Equal(t, 1, m.Increment("foo", 1))
	assert.Equal(t, 6, m.Increment("foo", 5))
	assert.Equal(t, 14, m.Increment("foo", 8))

	assert.Equal(t, 13, m.Decrement("foo", 1))
	assert.Equal(t, 8, m.Decrement("foo", 5))
	assert.Equal(t, 0, m.Decrement("foo", 8))

	// test overflow
	assert.True(t, m.Put("foo", "bar", time.Now().Add(time.Second*10)))
	assert.Equal(t, "bar", m.Get("foo").(string))
	assert.Equal(t, 1, m.Increment("foo", 1))
}

func TestMemoryStore_Forever(t *testing.T) {
	m := NewMemoryStore()

	assert.True(t, m.Forever("foo", "bar"))
	assert.Equal(t, "bar", m.Get("foo").(string))

	for _, r := range m.getRecords() {
		assert.Equal(t, time.Time{}, r.expired)
	}
}

func TestMemoryStore_Forget(t *testing.T) {
	m := NewMemoryStore()

	m.Put("foo", "bar", time.Now().Add(time.Second*1))
	assert.Equal(t, "bar", m.Get("foo").(string))
	assert.True(t, m.Forget("foo"))
	assert.Nil(t, m.Get("foo"))

	assert.False(t, m.Forget("foo2"))
}

func TestMemoryStore_Flush(t *testing.T) {
	m := NewMemoryStore()

	m.Put("foo", "bar", time.Now().Add(time.Second*1))
	m.Put("foo2", "bar2", time.Now().Add(time.Second*1))
	assert.Equal(t, 2, len(m.getRecords()))

	assert.True(t, m.Flush())
	assert.Equal(t, 0, len(m.getRecords()))
}

func BenchmarkMemoryStore_PutAndGet(b *testing.B) {
	m := NewMemoryStore()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// m.Put("foo", "bar", time.Now().Add(time.Second*1))
			// m.Get("foo")
			// m.Increment("foo", 1)
			m.Forget("foo")
		}
	})
}
