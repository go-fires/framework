package str

type String string

// String returns the string value.
func (s String) String() string {
	return string(s)
}

// Is checks if the string matches the pattern.
func (s String) Is(pattern string) bool {
	return Is(pattern, s.String())
}

// InArray checks if a string is in a string array.
func (s String) InArray(arr []string) bool {
	return InArray(s.String(), arr)
}

// Md5 returns the md5 hash of a string.
func (s String) Md5() string {
	return Md5(s.String())
}

// Sha1 returns the sha1 hash of a string.
func (s String) Sha1() string {
	return Sha1(s.String())
}

// Strpos returns the position of the first occurrence of a substring in a string.
func (s String) Strpos(substr string) int {
	return Strpos(s.String(), substr)
}

// Strrpos returns the position of the last occurrence of a substring in a string.
func (s String) Strrpos(substr string) int {
	return Strrpos(s.String(), substr)
}

// Strrev returns a string with the characters in reverse order.
func (s String) Strrev() string {
	return Strrev(s.String())
}

// Strtr returns a string with the characters in reverse order.
func (s String) Strtr(from, to string) string {
	return Strtr(s.String(), from, to)
}

// Shuffle returns a string with the characters in reverse order.
func (s String) Shuffle() string {
	return Shuffle(s.String())
}

// Atoi returns the integer value of a string.
func (s String) Atoi() AtoiRes {
	return Atoi(s.String())
}

// StrPad returns a string padded to a certain length with another string.
func (s String) StrPad(length int, pad string, padType StrPadType) string {
	return StrPad(s.String(), length, pad, padType)
}

// Bytes returns the bytes of the string.
func (s String) Bytes() []byte {
	return []byte(s)
}

// Length returns the length of the string.
func (s String) Length() int {
	return Length(s.String())
}

// Strcut returns part of a string.
func (s String) Strcut(start, length int) string {
	return Strcut(s.String(), start, length)
}

// Limit returns part of a string.
func (s String) Limit(length int, suffix ...string) string {
	return Limit(s.String(), length, suffix...)
}

// Ucfirst returns a string with the first character in uppercase.
func (s String) Ucfirst() string {
	return Ucfirst(s.String())
}

// Lcfirst returns a string with the first character in lowercase.
func (s String) Lcfirst() string {
	return Lcfirst(s.String())
}

// HtmlspecialcharsDecode returns a string with special HTML characters decoded.
func (s String) HtmlspecialcharsDecode() string {
	return HtmlspecialcharsDecode(s.String())
}

// Htmlspecialchars returns a string with special HTML characters encoded.
func (s String) Htmlspecialchars() string {
	return Htmlspecialchars(s.String())
}

// Trim returns a string with whitespace stripped from the beginning and end of the string.
func (s String) Trim() string {
	return Trim(s.String())
}

// IsUuid returns true if the string is a valid uuid.
func (s String) IsUuid() bool {
	return IsUuid(s.String())
}
