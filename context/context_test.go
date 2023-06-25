package context

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContext(t *testing.T) {
	Set("a", 1)

	a, err := Get("a")
	assert.Nil(t, err)
	assert.Equal(t, 1, a)
	b, err := Get("b")
	assert.NotNil(t, err)
	assert.Nil(t, b)

	assert.True(t, Has("a"))
	assert.False(t, Has("b"))

	assert.Equal(t, map[string]interface{}{
		"a": 1,
	}, All())

	Delete("a")
	assert.False(t, Has("a"))

	Set("a", 1)
	Clear()

	assert.False(t, Has("a"))
	assert.Equal(t, 0, len(All()))
}
