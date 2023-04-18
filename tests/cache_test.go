package tests

import (
	"github.com/go-fires/framework/cache"
	"github.com/go-fires/framework/config"
	"github.com/go-fires/framework/facade"
	"github.com/go-fires/framework/redis"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	app := createApplication()

	app.Register(config.NewProvider(app.Container))
	app.Register(redis.NewProvider(app.Container))
	app.Register(cache.NewProvider(app.Container))

	var foo string
	facade.Cache().Store("redis").Set("foo", "bar", time.Second*10)
	assert.Nil(t, facade.Cache().Store("redis").Get("foo", &foo))
	assert.Equal(t, "bar", foo)
}
