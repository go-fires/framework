package cache

import (
	"github.com/go-fires/framework/support/ints"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMemoryStore_PutGet(t *testing.T) {
	m := NewMemoryStore()

	assert.True(t, m.Put("foo", "bar", time.Second*1))
	var foo string
	assert.Nil(t, m.Get("foo", &foo))
	assert.Equal(t, "bar", foo)

	time.Sleep(time.Second * 2)

	var foo2 string
	assert.Error(t, m.Get("foo", &foo2))
	assert.Equal(t, "", foo2)

	type User struct {
		Name string
		Age  int
	}

	var user User
	assert.True(t, m.Put("user", User{"foo", 18}, time.Second*1))
	assert.Nil(t, m.Get("user", &user))
	assert.Equal(t, "foo", user.Name)
	assert.Equal(t, 18, user.Age)
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
	assert.True(t, m.Put("foo", "bar", time.Second*10))
	var foo string
	assert.Nil(t, m.Get("foo", &foo))
	assert.Equal(t, "bar", foo)
	assert.Equal(t, 1, m.Increment("foo", 1))
}

func TestMemoryStore_Forever(t *testing.T) {
	m := NewMemoryStore()

	assert.True(t, m.Forever("foo", "bar"))
	var foo string
	assert.Nil(t, m.Get("foo", &foo))
	assert.Equal(t, "bar", foo)

	m.records.Range(func(key, value interface{}) bool {
		assert.Equal(t, time.Time{}, value.(*record).expired)
		return true
	})
}

func TestMemoryStore_Forget(t *testing.T) {
	m := NewMemoryStore()

	var foo string
	m.Put("foo", "bar", time.Second*1)
	assert.Nil(t, m.Get("foo", &foo))
	assert.Equal(t, "bar", foo)
	assert.True(t, m.Forget("foo"))
	assert.False(t, m.Has("foo"))

	assert.False(t, m.Forget("foo2"))
}

func TestMemoryStore_Flush(t *testing.T) {
	m := NewMemoryStore()

	m.Put("foo", "bar", time.Second*1)
	m.Put("foo2", "bar2", time.Second*1)
	assert.Equal(t, 2, m.Length())

	assert.True(t, m.Flush())
	assert.Equal(t, 0, m.Length())
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
