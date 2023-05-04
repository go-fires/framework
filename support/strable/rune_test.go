package strable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunes_String(t *testing.T) {
	assert.Equal(t, "abc", Runes("abc").String())

	assert.Equal(t, "abc", Runes("abc").String())
	assert.Equal(t, "张三", Runes("张三").String())
}

func TestRunes_Bytes(t *testing.T) {
	assert.Equal(t, []byte("abc"), Runes("abc").Bytes())
	assert.Equal(t, []byte("张三"), Runes("张三").Bytes())
}

func TestRunes_Len(t *testing.T) {
	assert.Equal(t, 3, Runes("abc").Len())
	assert.Equal(t, 2, Runes("张三").Len())
}

func TestRunes_Cap(t *testing.T) {
	assert.Equal(t, 3, Runes("abc").Cap())
	assert.Equal(t, 2, Runes("张三").Cap())
}

func TestRunes_Empty(t *testing.T) {
	// is empty
	assert.True(t, Runes("").IsEmpty())
	assert.False(t, Runes("abc").IsEmpty())
	assert.False(t, Runes("张三").IsEmpty())

	// is not empty
	assert.False(t, Runes("").IsNotEmpty())
	assert.True(t, Runes("abc").IsNotEmpty())
	assert.True(t, Runes("张三").IsNotEmpty())
}

func TestRunes_Equal(t *testing.T) {
	// is equal
	assert.True(t, Runes("abc").IsEqual("abc"))
	assert.True(t, Runes("张三").IsEqual("张三"))
	assert.False(t, Runes("abc").IsEqual("张三"))

	// is not equal
	assert.False(t, Runes("abc").IsNotEqual("abc"))
	assert.False(t, Runes("张三").IsNotEqual("张三"))
	assert.True(t, Runes("abc").IsNotEqual("张三"))
}
