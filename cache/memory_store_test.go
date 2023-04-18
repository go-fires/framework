package cache

import (
	"github.com/go-fires/framework/support/ints"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func BenchmarkMemoryStore_Incr(b *testing.B) {
	m := NewMemoryStore()

	var foo int
	counter := &ints.Counter{}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Increment("foo", 1)
			counter.Inc(1)
		}
	})

	assert.Nil(b, m.Get("foo", &foo))
	assert.Equal(b, counter.Val(), foo)
}
