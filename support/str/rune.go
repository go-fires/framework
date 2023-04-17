package str

type Rune rune
type Runes []Rune

// String returns the string of the runes.
func (r Runes) String() string {
	return string(r)
}

// Bytes returns the bytes of the string.
func (r Runes) Bytes() []byte {
	return []byte(r.String())
}

// Len returns the length of the runes.
func (r Runes) Len() int {
	return len(r)
}

// Cap returns the capacity of the runes.
func (r Runes) Cap() int {
	return cap(r)
}

// IsEmpty returns true if the runes is empty.
func (r Runes) IsEmpty() bool {
	return r.Len() == 0
}

// IsNotEmpty returns true if the runes is not empty.
func (r Runes) IsNotEmpty() bool {
	return !r.IsEmpty()
}

// IsEqual returns true if the runes is equal to the string.
func (r Runes) IsEqual(s string) bool {
	return r.String() == s
}

// IsNotEqual returns true if the runes is not equal to the string.
func (r Runes) IsNotEqual(s string) bool {
	return !r.IsEqual(s)
}

// Substr returns the substring of the runes.
func (r Runes) Substr(start, length int) string {
	if r.Len() <= start {
		return ""
	}

	if r.Len() <= start+length {
		return r[start:].String()
	}

	return r[start : start+length].String()
}
