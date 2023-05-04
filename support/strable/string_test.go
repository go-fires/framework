package strable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString_Is(t *testing.T) {
	assert.True(t, String("abc").Is("ab*"))
}

func TestString_InArray(t *testing.T) {
	assert.True(t, String("abc").InArray([]string{"abc", "def"}))
}

func TestString_Md5(t *testing.T) {
	assert.Equal(t, "900150983cd24fb0d6963f7d28e17f72", String("abc").Md5())
}

func TestString_Sha1(t *testing.T) {
	assert.Equal(t, "a9993e364706816aba3e25717850c26c9cd0d89d", String("abc").Sha1())
}

func TestString_Strpos(t *testing.T) {
	assert.Equal(t, 0, String("aabbcc").Strpos("a"))
}

func TestString_Strrpos(t *testing.T) {
	assert.Equal(t, 1, String("aabbcc").Strrpos("a"))
}

func TestString_Strrev(t *testing.T) {
	assert.Equal(t, "cba", String("abc").Strrev())
}

func TestString_Strtr(t *testing.T) {
	assert.Equal(t, "bbbbcc", String("aabbcc").Strtr("a", "b"))
}

func TestString_StrShuffle(t *testing.T) {
	assert.True(t, String("abc").InArray([]string{"abc", "acb", "bac", "bca", "cab", "cba"}))
}

func TestString_Atoi(t *testing.T) {
	assert.Equal(t, 1, String("1").Atoi().Val())
	assert.Equal(t, 2, String("2").Atoi().Val())
	assert.Error(t, String("a").Atoi().Err())
	assert.Error(t, String("1a").Atoi().Err())
	assert.True(t, String("1").Atoi().IsOk())
	assert.False(t, String("a").Atoi().IsOk())
	assert.False(t, String("1").Atoi().IsErr())
	assert.True(t, String("a").Atoi().IsErr())
}

func TestString_Bytes(t *testing.T) {
	assert.Equal(t, []byte("abc"), String("abc").Bytes())
}

func TestString_StrPad(t *testing.T) {
	assert.Equal(t, "   abc", String("abc").StrPad(6, " ", StrPadLeft))
	assert.Equal(t, "abc   ", String("abc").StrPad(6, " ", StrPadRight))
	assert.Equal(t, " abc  ", String("abc").StrPad(6, " ", StrPadBoth))
}

func TestString_Strcut(t *testing.T) {
	assert.Equal(t, "a", String("abc").Strcut(0, 1))
	assert.Equal(t, "ab", String("abc").Strcut(0, 2))
	assert.Equal(t, "bc", String("abc").Strcut(1, 2))
	assert.Equal(t, "c", String("abc").Strcut(2, 1))
}

func TestString_Limit(t *testing.T) {
	assert.Equal(t, "a...", String("abc").Limit(1, "..."))
	assert.Equal(t, "ab...", String("abc").Limit(2, "..."))
	assert.Equal(t, "abc", String("abc").Limit(3, "..."))
	assert.Equal(t, "abc", String("abc").Limit(4, "..."))
}

func TestString_Length(t *testing.T) {
	assert.Equal(t, 3, String("abc").Length())
}

func TestString_Ucfirst(t *testing.T) {
	assert.Equal(t, "Abc", String("abc").Ucfirst())
}

func TestString_Lcfirst(t *testing.T) {
	assert.Equal(t, "abc", String("Abc").Lcfirst())
}

func TestString_HTml(t *testing.T) {
	assert.Equal(t, "a&lt;b&gt;c", String("a<b>c").Htmlspecialchars())
	assert.Equal(t, "a<b>c", String("a&lt;b&gt;c").HtmlspecialcharsDecode())
}

func TestString_Trim(t *testing.T) {
	assert.Equal(t, "abc", String(" abc ").Trim())
	assert.Equal(t, "abc", String("\r abc \n\r").Trim())
}

func TestString_IsUuid(t *testing.T) {
	assert.True(t, String("b2c6b2d0-0dc2-11e8-8eb2-f2801f1b9fd1").IsUuid())
	assert.False(t, String("b2c6b2d0-0dc2-11e8-8eb2-f2801f1b9fd").IsUuid())
}
