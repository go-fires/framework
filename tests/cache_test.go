package tests

import (
	"testing"
	"time"

	"github.com/go-fires/fires/facade"
	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	createApplication()

	var foo string
	facade.Cache().Store("redis").Set("foo", "bar", time.Second*10)
	assert.Nil(t, facade.Cache().Store("redis").Get("foo", &foo))
	assert.Equal(t, "bar", foo)
}
