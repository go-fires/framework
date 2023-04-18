package strs

import "strconv"

// AtoiRes is the result of Atoi.
type AtoiRes struct {
	val int
	err error
}

// Val returns the value of the result.
func (r AtoiRes) Val() int {
	return r.val
}

// Err returns the error of the result.
func (r AtoiRes) Err() error {
	return r.err
}

// IsErr returns true if the result is an error.
func (r AtoiRes) IsErr() bool {
	return r.err != nil
}

// IsOk returns true if the result is ok.
func (r AtoiRes) IsOk() bool {
	return r.err == nil
}

// Atoi converts a string to an int.
func Atoi(s string) AtoiRes {
	val, err := strconv.Atoi(s)

	return AtoiRes{
		val: val,
		err: err,
	}
}
