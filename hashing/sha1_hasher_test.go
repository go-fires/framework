package hashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSha1Hasher(t *testing.T) {
	value := "hello"
	hashedValue, err := Sha1.Make(value)

	assert.Nil(t, err)
	assert.Equal(t, "68656c6c6fda39a3ee5e6b4b0d3255bfef95601890afd80709", hashedValue)
	assert.True(t, Sha1.Check(value, hashedValue))

	assert.True(t, Sha1.Check(value, Sha1.MustMake(value)))
}
