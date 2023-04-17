package hashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSha1Hasher(t *testing.T) {
	hasher := NewSha1Hasher()

	value := "hello"
	hashedValue, err := hasher.Make(value)

	assert.Nil(t, err)
	assert.Equal(t, "68656c6c6fda39a3ee5e6b4b0d3255bfef95601890afd80709", hashedValue)
	assert.True(t, hasher.Check(value, hashedValue))
}
