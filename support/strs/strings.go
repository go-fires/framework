package strs

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"regexp"
	"strings"
	"unicode"
)

// Is returns true if the value matches the pattern.
// The pattern can contain the wildcard character *.
//
// Example:
//
//	Is("*.example.com", "www.example.com") // true
//	Is("*.example.com", "example.com") // false
func Is(pattern, value string) bool {
	if pattern == value {
		return true
	}

	pattern = strings.ReplaceAll(pattern, "*", ".*")

	match, err := regexp.Match(pattern, []byte(value))
	if err != nil {
		return false
	}

	return match
}

// InArray checks if a string is in a string array.
//
// Example:
//
//	InArray("abc", []string{"abc", "def"}) // true
func InArray(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

// Md5 returns the md5 hash of a string.
//
// Example:
//
//	Md5("abc") // 900150983cd24fb0d6963f7d28e17f72
func Md5(s string) string {
	sm := md5.Sum([]byte(s))

	return fmt.Sprintf("%x", sm)
}

// Sha1 returns the sha1 hash of a string.
//
// Example:
//
//	Sha1("abc") // a9993e364706816aba3e25717850c26c9cd0d89d
func Sha1(s string) string {
	sm := sha1.Sum([]byte(s))

	return fmt.Sprintf("%x", sm)
}

// Strpos returns the position of the first occurrence of a substring in a string.
//
// Example:
//
//	Strpos("abc", "b") // 1
func Strpos(haystack, needle string) int {
	return strings.Index(haystack, needle)
}

// Strrpos returns the position of the last occurrence of a substring in a string.
//
// Example:
//
//	Strrpos("abc", "b") // 1
func Strrpos(haystack, needle string) int {
	return strings.LastIndex(haystack, needle)
}

// Strrev returns a string with its characters in reverse order.
//
// Example:
//
//	Strrev("abc") // "cba"
func Strrev(s string) string {
	var reversed string

	for _, v := range s {
		reversed = string(v) + reversed
	}

	return reversed
}

// Strtr replaces all occurrences of the search string with the replacement string.
//
// Example:
//
//	Strtr("aabbcc", "a", "b") // "bbbbcc"
func Strtr(s, from, to string) string {
	return strings.NewReplacer(from, to).Replace(s)
}

// StrtrArray replaces all occurrences of the search strings with the replacement strings.
//
// Example:
//
//	StrtrArray("aabbcc", map[string]string{"a": "d", "b": "f"} // "ddffcc"
//	StrtrArray("aabbcc", map[string]string{"a": "b", "b": "f"}) // "ffffcc" (wrong) todo: to be fixed
// func StrtrArray(s string, rp map[string]string) string {
//
// 	strings.ReplaceAll(s, "a", "b")
// 	for k, v := range rp {
// 		s = strings.ReplaceAll(s, k, v)
// 	}
//
// 	return s
// }

// Shuffle returns a string with its characters in random order.
//
// Example:
//
//	Shuffle("abc") // "bca"
func Shuffle(s string) string {
	ss := strings.Split(s, "")

	rand.Shuffle(len(ss), func(i, j int) {
		ss[i], ss[j] = ss[j], ss[i]
	})

	return strings.Join(ss, "")
}

// randomLetters is the letters used in RandomString.
const randomLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var (
	randomLettersRunes       = []rune(randomLetters)
	randomLettersRunesLength = len(randomLetters)
)

// RandomString returns a random string with the specified length.
//
// Example:
//
//	RandomString(10) // "qujrlkhyqr"
func RandomString(length int) string {
	b := make([]rune, length)

	for i := range b {
		b[i] = randomLettersRunes[rand.Intn(randomLettersRunesLength)]
	}

	return string(b)
}

// StrPadType is the type of padding.
type StrPadType int

// StrPadType constants.
const (
	StrPadLeft StrPadType = iota
	StrPadRight
	StrPadBoth
)

// StrPad returns an input string padded on the left or right to specified length with pad string.
//
// Example:
//
//	StrPad("abc", 6, " ", StrPadLeft) // "   abc"
//	StrPad("abc", 6, " ", StrPadRight) // "abc   "
//	StrPad("abc", 6, " ", StrPadBoth) // " abc  "
func StrPad(s string, length int, pad string, padType StrPadType) string {
	if len(s) >= length {
		return s
	}

	switch padType {
	case StrPadLeft:
		return strings.Repeat(pad, length-len(s)) + s
	case StrPadRight:
		return s + strings.Repeat(pad, length-len(s))
	case StrPadBoth:
		length = length - len(s)
		paddingLeft := strings.Repeat(pad, (length)/2)
		paddingRight := strings.Repeat(pad, (length)/2+(length)%2)
		return paddingLeft + s + paddingRight
	default:
		return s
	}
}

// Length returns the length of a string.
// Support chinese characters.
//
// Example:
//
//	Length("abc") // 3
//	Length("张三李四") // 4
func Length(s string) int {
	return Runes(s).Len()
}

// Len is alias of Length.
func Len(s string) int {
	return Length(s)
}

// Strcut returns a string with a specified length starting from a specified position.
// Support chinese characters.
//
// Example:
//
//	Strcut("abc", 1, 1) // "b"
//	Strcut("张三李四", 1, 2) // "三李
func Strcut(s string, start, length int) string {
	sr := []rune(s)

	if len(sr) <= start {
		return ""
	}

	if len(sr) <= start+length {
		return string(sr[start:])
	}

	return string(sr[start : start+length])
}

// Limit returns a string with a specified length.
// Support chinese characters.
//
// Example:
//
//	Limit("abc", 2) // "ab"
//	Limit("张三李四", 2) // "张三"
//	Limit("张三李四", 2, "...") // "张三..."
func Limit(s string, length int, suffix ...string) string {
	r, sf := Runes(s), ""

	if r.Len() <= length {
		return s
	}

	if len(suffix) > 0 {
		sf = suffix[0]
	}

	return r.Substr(0, length) + sf
}

// Ucfirst returns a string with the first character in uppercase.
func Ucfirst(str string) string {
	r := Runes(str)

	if r.Len() < 1 {
		return str
	}

	return strings.ToUpper(r.Substr(0, 1)) + r.Substr(1, r.Len())
}

// Lcfirst returns a string with the first character in lowercase.
func Lcfirst(str string) string {
	r := Runes(str)

	if r.Len() < 1 {
		return str
	}

	return strings.ToLower(r.Substr(0, 1)) + r.Substr(1, r.Len())
}

// ReplaceVars replaces all occurrences of the search strings with the replacement strings.
func ReplaceVars(str string, vars map[string]string) string {
	for key, value := range vars {
		str = strings.ReplaceAll(str, "{"+key+"}", value)
	}

	return str
}

// HtmlspecialcharsDecode converts special HTML entities back to characters.
func HtmlspecialcharsDecode(str string) string {
	str = strings.ReplaceAll(str, "&amp;", "&")
	str = strings.ReplaceAll(str, "&lt;", "<")
	str = strings.ReplaceAll(str, "&gt;", ">")
	str = strings.ReplaceAll(str, "&quot;", "\"")
	str = strings.ReplaceAll(str, "&#039;", "'")

	return str
}

// Htmlspecialchars converts special characters to HTML entities.
func Htmlspecialchars(str string) string {
	str = strings.ReplaceAll(str, "&", "&amp;")
	str = strings.ReplaceAll(str, "<", "&lt;")
	str = strings.ReplaceAll(str, ">", "&gt;")
	str = strings.ReplaceAll(str, "\"", "&quot;")
	str = strings.ReplaceAll(str, "'", "&#039;")

	return str
}

// Trim returns a string with space stripped from the beginning and end of str.
func Trim(str string) string {
	return strings.TrimSpace(str)
}

// IsUUID returns true if the string is a valid UUID.
func IsUUID(str string) bool {
	pattern := "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	match, err := regexp.MatchString(pattern, str)

	if err != nil {
		return false
	}

	return match
}

// UUID returns a UUID.
func UUID() string {
	return uuid.New().String()
}

// ReplaceLast replaces the last occurrence of a string.
//
// Example:
//
//	ReplaceLast("abc", "c", "d") // "abd"
//	ReplaceLast("abc", "d", "e") // "abc"
func ReplaceLast(str, old, new string) string {
	if old == "" {
		return str
	}

	if pos := strings.LastIndex(str, old); pos == -1 {
		return str
	} else {
		return str[:pos] + new + str[pos+len(old):]
	}
}

// LowerFirst returns a string with the first character in uppercase.
func LowerFirst(str string) string {
	rs := []rune(str)

	if len(rs) < 1 {
		return str
	}

	rs[0] = unicode.ToLower(rs[0])

	return string(rs)
}

// UpperFirst returns a string with the first character in uppercase.
func UpperFirst(str string) string {
	rs := []rune(str)

	if len(rs) < 1 {
		return str
	}

	rs[0] = unicode.ToUpper(rs[0])

	return string(rs)
}
