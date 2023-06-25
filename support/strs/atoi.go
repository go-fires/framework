package strs

import "strconv"

// AtoiResult is the result of Atoi.
type AtoiResult struct {
	val int
	err error
}

// Val returns the value of the result.
func (r AtoiResult) Val() int {
	return r.val
}

// Err returns the error of the result.
func (r AtoiResult) Err() error {
	return r.err
}

// IsErr returns true if the result is an error.
func (r AtoiResult) IsErr() bool {
	return r.err != nil
}

// IsOk returns true if the result is ok.
func (r AtoiResult) IsOk() bool {
	return r.err == nil
}

// Atoi converts a string to an int.
func Atoi(s string) AtoiResult {
	val, err := strconv.Atoi(s)

	return AtoiResult{
		val: val,
		err: err,
	}
}
