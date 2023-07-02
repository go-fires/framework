package hashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5Hasher(t *testing.T) {
	value := "123456"
	hashedValue, err := Md5.Make(value)

	assert.Nil(t, err)
	assert.True(t, Md5.Check(value, hashedValue))

	assert.True(t, Md5.Check(value, Md5.MustMake(value)))
}
