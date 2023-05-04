package strable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
	assert.Equal(t, 1, Atoi("1").Val())
	assert.Equal(t, 2, Atoi("2").Val())
	assert.Error(t, Atoi("a").Err())
	assert.Error(t, Atoi("1a").Err())
	assert.True(t, Atoi("1").IsOk())
	assert.False(t, Atoi("a").IsOk())
	assert.False(t, Atoi("1").IsErr())
	assert.True(t, Atoi("a").IsErr())
}
