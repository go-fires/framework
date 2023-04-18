package cache

import (
	"github.com/go-fires/framework/support/ints"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func createRedisStore() *RedisStore {
	r := NewRedisStore(redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	}))

	r.Flush()

	return r
}

func TestRedisStore_PutGet(t *testing.T) {
	m := createRedisStore()

	assert.True(t, m.Put("foo", "bar", time.Second*1))
	assert.Equal(t, "bar", m.Get("foo").(string))

	time.Sleep(time.Second * 2)
	assert.Nil(t, m.Get("foo"))

	type User struct {
		Name string
		Age  int
	}

	assert.True(t, m.Put("user", User{"foo", 18}, time.Second*1))
	// fmt.Println(m.Get("user"))
	// assert.Equal(t, User{"foo", 18}, m.Get("user").(User))
	// assert.Equal(t, "foo", m.Get("user").(User).Name)
}

func TestRedisStore_IncrAndDecr(t *testing.T) {
	m := createRedisStore()

	assert.Equal(t, 1, m.Increment("foo", 1))
	assert.Equal(t, 6, m.Increment("foo", 5))
	assert.Equal(t, 14, m.Increment("foo", 8))

	assert.Equal(t, 13, m.Decrement("foo", 1))
	assert.Equal(t, 8, m.Decrement("foo", 5))
	assert.Equal(t, 0, m.Decrement("foo", 8))

	// test overflow
	assert.True(t, m.Put("foo", "bar", time.Second*10))
	assert.Equal(t, "bar", m.Get("foo").(string))
	assert.Equal(t, 1, m.Increment("foo", 1))
}

// func TestRedisStore_Forever(t *testing.T) {
// 	m := createRedisStore()
//
// 	assert.True(t, m.Forever("foo", "bar"))
// 	assert.Equal(t, "bar", m.Get("foo").(string))
//
// 	m.records.Range(func(key, value interface{}) bool {
// 		assert.Equal(t, time.Time{}, value.(*record).expired)
// 		return true
// 	})
// }

func TestRedisStore_Forget(t *testing.T) {
	m := createRedisStore()

	m.Put("foo", "bar", time.Second*1)
	assert.Equal(t, "bar", m.Get("foo").(string))
	assert.True(t, m.Forget("foo"))
	assert.Nil(t, m.Get("foo"))

	assert.False(t, m.Forget("foo2"))
}

// func TestRedisStore_Flush(t *testing.T) {
// 	m := createRedisStore()
//
// 	m.Put("foo", "bar", time.Second*1)
// 	m.Put("foo2", "bar2", time.Second*1)
// 	assert.Equal(t, 2, m.Length())
//
// 	assert.True(t, m.Flush())
// 	assert.Equal(t, 0, m.Length())
// }

func BenchmarkRedisStore_PutAndGet(b *testing.B) {
	m := createRedisStore()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// m.Put("foo", "bar", time.Now().Add(time.Second*1))
			// m.Get("foo")
			// m.Increment("foo", 1)
			m.Forget("foo")
		}
	})
}

func BenchmarkRedisStore_Incr(b *testing.B) {
	m := createRedisStore()

	counter := &ints.Counter{}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Increment("foo", 1)
			counter.Inc(1)
		}
	})

	assert.Equal(b, counter.Val(), m.Get("foo"))
}
