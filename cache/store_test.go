package cache

import (
	"github.com/go-fires/framework/contracts/cache"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func createMemoryStore() *MemoryStore {
	return NewMemoryStore()
}

func createRedisStore() *RedisStore {
	r := NewRedisStore(redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	}))

	r.Flush()

	return r
}

var tests = []struct {
	name  string
	store cache.Store
}{
	{"redis", createMemoryStore()},
	{"memory", createRedisStore()},
}

func TestStore_Putget(t *testing.T) {
	for _, tt := range tests {
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

func testIncrAndDecr(store cache.Store, t *testing.T) {
	assert.Equal(t, 1, store.Increment("foo", 1))
	assert.Equal(t, 6, store.Increment("foo", 5))
	assert.Equal(t, 14, store.Increment("foo", 8))

	assert.Equal(t, 13, store.Decrement("foo", 1))
	assert.Equal(t, 8, store.Decrement("foo", 5))
	assert.Equal(t, 0, store.Decrement("foo", 8))

	// test overflow
	assert.True(t, store.Put("foo", "bar", time.Second*10))
	var foo string
	assert.Nil(t, store.Get("foo", &foo))
	assert.Equal(t, "bar", foo)
	assert.Equal(t, 0, store.Increment("foo", 1))
}

func testForever(store cache.Store, t *testing.T) {
	assert.True(t, store.Forever("foo", "bar"))
	var foo string
	assert.Nil(t, store.Get("foo", &foo))
	assert.Equal(t, "bar", foo)
}

func testForget(store cache.Store, t *testing.T) {
	var foo string
	store.Put("foo", "bar", time.Second*1)
	assert.Nil(t, store.Get("foo", &foo))
	assert.Equal(t, "bar", foo)
	assert.True(t, store.Forget("foo"))
	assert.False(t, store.Has("foo"))

	assert.False(t, store.Forget("foo2"))
}

func testFlush(store cache.Store, t *testing.T) {
	var foo, foo2 string
	store.Put("foo", "bar", time.Second*1)
	store.Put("foo2", "bar2", time.Second*1)
	assert.Nil(t, store.Get("foo", &foo))
	assert.Nil(t, store.Get("foo2", &foo2))

	assert.True(t, store.Flush())
	var foo3, foo4 string
	assert.Error(t, store.Get("foo", &foo3))
	assert.Error(t, store.Get("foo2", &foo4))
}

func TestStore_Memory(t *testing.T) {
	testIncrAndDecr(createMemoryStore(), t)
	testForever(createMemoryStore(), t)
	testForget(createMemoryStore(), t)
	testFlush(createMemoryStore(), t)
}

func TestStore_Redis(t *testing.T) {
	testIncrAndDecr(createRedisStore(), t)
	testForever(createRedisStore(), t)
	testForget(createRedisStore(), t)
	testFlush(createRedisStore(), t)
}
