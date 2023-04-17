package hashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBcryptHasher(t *testing.T) {
	value := "123456"

	hash := NewBcryptHasher()
	hashedValue, _ := hash.Make(value)
	assert.True(t, hash.Check(value, hashedValue))

	hash2 := &BcryptHasher{}
	hashedValue2 := hash2.MustMake(value)
	assert.True(t, hash2.Check(value, hashedValue2))
}
