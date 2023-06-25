package md5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5Hasher(t *testing.T) {
	value := "123456"
	hashedValue, err := Make(value)

	assert.Nil(t, err)
	assert.True(t, Check(value, hashedValue))

	assert.True(t, Check(value, MustMake(value)))
}
