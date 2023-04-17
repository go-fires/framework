package tests

import (
	"context"
	"github.com/go-fires/framework/facade"
	"github.com/go-fires/framework/redis"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	app := createApplication()
	var ctx = context.Background()

	app.Register(redis.NewProvider(app.Container))

	facade.Redis().Connect().Set(ctx, "key", "value", time.Second*5)
	assert.Equal(t, "value", facade.Redis().Connect().Get(ctx, "key").Val())
}
