package strs

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"unicode"
)

func TestIs(t *testing.T) {
	assert.True(t, Is("abc", "abc"))
	assert.False(t, Is("abcc", "abc"))
	assert.True(t, Is("ab*", "abc"))
	assert.True(t, Is("ab*", "ab"))
	assert.True(t, Is("ab/*", "ab/cc"))
	assert.True(t, Is("ab/*", "ab/"))
	assert.True(t, Is("*", "ab"))
	assert.True(t, Is("*", ""))
	assert.True(t, Is("ab/*", "ab/cc/dd"))
	assert.True(t, Is("*dd/", "ab/cc/dd/"))
	assert.False(t, Is("*dd/d", "dd/"))
}

func TestInArray(t *testing.T) {
	assert.True(t, InArray("1", []string{"1", "2"}))
	assert.True(t, InArray("2", []string{"1", "2"}))
	assert.False(t, InArray("3", []string{"1", "2"}))
	assert.False(t, InArray("12", []string{"1", "2"}))
}

func TestMd5(t *testing.T) {
	assert.Equal(t, "900150983cd24fb0d6963f7d28e17f72", Md5("abc"))
}

func TestSha1(t *testing.T) {
	assert.Equal(t, "a9993e364706816aba3e25717850c26c9cd0d89d", Sha1("abc"))
}

func TestStrpos(t *testing.T) {
	assert.Equal(t, 0, Strpos("aabbcc", "a"))
	assert.Equal(t, 2, Strpos("aabbcc", "b"))
	assert.Equal(t, -1, Strpos("aabbcc", "d"))
}

func TestStrrpos(t *testing.T) {
	assert.Equal(t, 1, Strrpos("aabbcc", "a"))
	assert.Equal(t, 3, Strrpos("aabbcc", "b"))
	assert.Equal(t, -1, Strrpos("aabbcc", "d"))
}

func TestStrrev(t *testing.T) {
	assert.Equal(t, "cba", Strrev("abc"))
}

func TestStrtr(t *testing.T) {
	assert.Equal(t, "bbbbcc", Strtr("aabbcc", "a", "b"))
}

// todo: to be fixed
// func TestStrtrArray(t *testing.T) {
// 	assert.Equal(t, "ddffcc", StrtrArray("aabbcc", map[string]string{"a": "d", "b": "f"}))
// 	assert.Equal(t, "ffffcc", StrtrArray("aabbcc", map[string]string{"a": "b", "b": "f"}))
// }

func TestStrShuffle(t *testing.T) {
	assert.True(t, InArray(Shuffle("abc"), []string{"abc", "acb", "bac", "bca", "cab", "cba"}))
}

func TestRandomString(t *testing.T) {
	r1, r2 := RandomString(10), RandomString(10)

	assert.Equal(t, 10, len(r1))
	assert.Equal(t, 10, len(r2))
	assert.NotEqual(t, r1, r2)
}

func TestStrPad(t *testing.T) {
	assert.Equal(t, "---abc", StrPad("abc", 6, "-", StrPadLeft))
	assert.Equal(t, "abc---", StrPad("abc", 6, "-", StrPadRight))
	assert.Equal(t, "--abc--", StrPad("abc", 7, "-", StrPadBoth))
	assert.Equal(t, "abc", StrPad("abc", 2, "-", StrPadBoth))
	assert.Equal(t, "--abc---", StrPad("abc", 8, "-", StrPadBoth))
	assert.Equal(t, "abc", StrPad("abc", 2, "-", StrPadBoth))
}

func TestLength(t *testing.T) {
	assert.Equal(t, 3, Length("abc"))
	assert.Equal(t, 4, Length("张三李四"))

	// len alias `Length`
	assert.Equal(t, 3, Len("abc"))
	assert.Equal(t, 4, Len("张三李四"))
}

func TestStrcut(t *testing.T) {
	assert.Equal(t, "abc", Strcut("abc", 0, 3))
	assert.Equal(t, "张三李", Strcut("张三李四", 0, 3))
	assert.Equal(t, "张三李四", Strcut("张三李四", 0, 4))
	assert.Equal(t, "张三李四", Strcut("张三李四", 0, 5))
	assert.Equal(t, "张三李四", Strcut("张三李四", 0, 6))
	assert.Equal(t, "李四", Strcut("张三李四", 2, 2))
}

func TestLimit(t *testing.T) {
	assert.Equal(t, "abc", Limit("abc", 3))
	assert.Equal(t, "ab", Limit("abc", 2))
	assert.Equal(t, "a", Limit("abc", 1))
	assert.Equal(t, "", Limit("abc", 0))
	assert.Equal(t, "张三", Limit("张三李四", 2))
	assert.Equal(t, "张...", Limit("张三李四", 1, "..."))
}

func TestUcfirst(t *testing.T) {
	assert.Equal(t, "Abc", Ucfirst("abc"))
	assert.Equal(t, "张三", Ucfirst("张三"))
}

func TestLcfirst(t *testing.T) {
	assert.Equal(t, "abc", Lcfirst("Abc"))
	assert.Equal(t, "张三", Lcfirst("张三"))
}

func TestReplaceVars(t *testing.T) {
	assert.Equal(t, "adc", ReplaceVars("a{b}c", map[string]string{"b": "d"}))
	assert.Equal(t, "adc", ReplaceVars("a{b}c", map[string]string{"b": "d", "c": "e"}))
}

func TestHtml(t *testing.T) {
	assert.Equal(t, "a&lt;b&gt;c", Htmlspecialchars("a<b>c"))
	assert.Equal(t, "a<b>c", HtmlspecialcharsDecode("a&lt;b&gt;c"))
}

func TestTrim(t *testing.T) {
	assert.Equal(t, "abc", Trim(" abc "))
	assert.Equal(t, "abc", Trim("abc"))
	assert.Equal(t, "abc", Trim("\r abc \n"))
}

func TestIsUuid(t *testing.T) {
	assert.True(t, IsUUID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6"))
	assert.False(t, IsUUID("f81d4fae-7dec-11d0-a765-00a0c91e6bf"))
	assert.False(t, IsUUID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6a"))
}

func TestUuid(t *testing.T) {
	assert.Equal(t, 36, len(UUID()))
	assert.True(t, IsUUID(UUID()))
}

func TestReplaceLast(t *testing.T) {
	assert.Equal(t, "foobar fooqux", ReplaceLast("foobar foobar", "bar", "qux"))
	assert.Equal(t, "foo/bar? foo/qux?", ReplaceLast("foo/bar? foo/bar?", "bar?", "qux?"))
	assert.Equal(t, "foobar foo", ReplaceLast("foobar foobar", "bar", ""))
	assert.Equal(t, "foobar foobar", ReplaceLast("foobar foobar", "xxx", "yyy"))
	assert.Equal(t, "foobar foobar", ReplaceLast("foobar foobar", "", "yyy"))
	assert.Equal(t, "Malmö Jönkxxxping", ReplaceLast("Malmö Jönköping", "ö", "xxx"))
	assert.Equal(t, "Malmö Jönköping", ReplaceLast("Malmö Jönköping", "", "yyy"))
	assert.Equal(t, "我是yyy国人", ReplaceLast("我是中国人", "中", "yyy"))
	assert.Equal(t, "我是中国人", ReplaceLast("我是中国人", "", "yyy"))
}

func TestSSS(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"", "FiresGOComponent", "fires_g_o_component"},
		{"", "FiresGoComponent", "fires_go_component"},
		{"", "FiresGoComponent", "fires_go_component"},
		{"", "Fires Go Component", "fires_go_component"},
		{"", "Fires    Go      Component   ", "fires_go_component"},
		{"", "FiresGoComponent", "fires__go__component"},
		{"", "FiresGoComponent_", "fires_go_component_"},
		{"", "fires go Component", "fires_go_component"},
		{"", "fires go MoreComponent", "fires_go_more_component"},
		{"", "foo-bar", "foo-bar"},
		{"", "Foo-Bar", "foo-_bar"},
		{"", "Foo_Bar", "foo__bar"},
		{"", "ŻółtaŁódka", "żółtałódka"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.Map(func(r rune) rune {
				if unicode.IsSpace(r) || r == '-' || r == '_' {
					return -1
				}
				if unicode.IsLower(r) {
					return unicode.ToUpper(r)
				}
				return r
			}, tt.args); got != tt.want {
				t.Errorf("Snake(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}

func TestLowerFirst(t *testing.T) {
	assert.Equal(t, "foo", LowerFirst("Foo"))
	assert.Equal(t, "foo", LowerFirst("foo"))
	assert.Equal(t, "f", LowerFirst("F"))
	assert.Equal(t, "", LowerFirst(""))
	assert.Equal(t, "1", LowerFirst("1"))
	assert.Equal(t, "中国", LowerFirst("中国"))
	assert.Equal(t, "żółtałódka", LowerFirst("żółtałódka"))
}
