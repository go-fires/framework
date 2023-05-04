package strable

// StrPadType is the type of padding.
type StrPadType int

// StrPadType constants.
const (
	StrPadLeft StrPadType = iota
	StrPadRight
	StrPadBoth
)

// randomStringLetters is the letters used in RandomString.
const randomStringLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
