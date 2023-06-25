package bcrypt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBcryptHasher(t *testing.T) {
	value := "123456"

	hashedValue, _ := Make(value)
	assert.True(t, Check(value, hashedValue))

	hashedValue2 := MustMake(value)
	assert.True(t, Check(value, hashedValue2))
}
