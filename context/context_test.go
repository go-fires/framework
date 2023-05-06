package context

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContext(t *testing.T) {
	assert.Nil(t, Get("test"))

	Set("test", "test")
	assert.Equal(t, "test", Get("test"))

	assert.True(t, Has("test"))
	assert.False(t, Has("test2"))

	Delete("test")
	assert.False(t, Has("test"))

	Set("test", "test")
	assert.True(t, Has("test"))
	Clear()
	assert.False(t, Has("test"))
}
