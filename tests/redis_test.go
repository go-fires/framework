package tests

import (
	"context"
	"testing"
	"time"

	"github.com/go-fires/framework/facade"
	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	createApplication()
	var ctx = context.Background()

	facade.Redis().Connect().Set(ctx, "key", "value", time.Second*5)
	assert.Equal(t, "value", facade.Redis().Connect().Get(ctx, "key").Val())
}
