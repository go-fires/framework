package hashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBcryptHasher(t *testing.T) {
	value := "123456"

	hashedValue, _ := Bcrypt.Make(value)
	assert.True(t, Bcrypt.Check(value, hashedValue))

	hashedValue2 := Bcrypt.MustMake(value)
	assert.True(t, Bcrypt.Check(value, hashedValue2))
}
