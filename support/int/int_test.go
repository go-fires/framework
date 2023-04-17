package int

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt_InArray(t *testing.T) {
	assert.True(t, Int(1).InArray([]int{1, 2, 3}))
	assert.False(t, Int(4).InArray([]int{1, 2, 3}))
}

func TestInt_Itoa(t *testing.T) {
	assert.Equal(t, "1", Int(1).Itoa())
	assert.Equal(t, "2", Int(2).Itoa())
}

func TestInt_String(t *testing.T) {
	assert.Equal(t, "1", Int(1).String())
	assert.Equal(t, "2", Int(2).String())
}

func TestInt_Bytes(t *testing.T) {
	assert.Equal(t, []byte("1"), Int(1).Bytes())
	assert.Equal(t, []byte("2"), Int(2).Bytes())
}

func TestInt_Val(t *testing.T) {
	assert.Equal(t, 1, Int(1).Val())
	assert.Equal(t, 2, Int(2).Val())
}

func TestInt_Between(t *testing.T) {
	assert.True(t, Int(1).Between(1, 2))
	assert.True(t, Int(2).Between(1, 2))
	assert.False(t, Int(3).Between(1, 2))
}
