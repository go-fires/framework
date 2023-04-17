package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var ctx = context.Background()

func TestRediser(t *testing.T) {
	m := New(&Config{
		Default: "db1",
		Connections: map[string]Configable{
			"db1": &redis.Options{
				Addr:     "localhost:6379",
				Password: "", // no password set
				DB:       0,  // use default DB
			},
			"db2": &redis.Options{
				Addr:     "localhost:6379",
				Password: "", // no password set
				DB:       1,  // use default DB
			},
			"db3": &redis.ClusterOptions{
				Addrs: []string{"localhost:6379", "localhost:6379"},
			},
		},
	})

	// use default(db1)
	assert.Nil(t, m.Connect().Set(ctx, "aaa", "1", time.Second*2).Err())
	assert.Equal(t, "1", m.Connect().Get(ctx, "aaa").Val())

	// db1
	assert.Equal(t, "1", m.Connect("db1").Get(ctx, "aaa").Val())
	assert.Error(t, m.Connect("db2").Get(ctx, "aaa").Err())

	// db2
	assert.Nil(t, m.Connect("db2").Set(ctx, "bbb", "1", time.Second*2).Err())
	assert.Equal(t, "1", m.Connect("db2").Get(ctx, "bbb").Val())

	// test sleep
	time.Sleep(time.Second * 3)
	assert.Equal(t, "", m.Connect().Get(ctx, "aaa").Val())

	// not found
	assert.Panics(t, func() {
		m.Connect("not-found").Get(ctx, "aaa").Val()
	})

	// type
	assert.Equal(t, "localhost:6379", m.Connect("db1").(*redis.Client).Options().Addr)
}

func BenchmarkMap(b *testing.B) {
	m := New(&Config{
		Default: "db1",
		Connections: map[string]Configable{
			"db1": &redis.Options{
				Addr:     "localhost:6379",
				Password: "", // no password set
				DB:       0,  // use default DB
			},
			"db2": &redis.Options{
				Addr:     "localhost:6379",
				Password: "", // no password set
				DB:       1,  // use default DB
			},
		},
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			go m.Connect().Get(ctx, "default").Val()
		}
	})
}
