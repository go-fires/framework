package strable

import "strconv"

// AtoiResp is the result of Atoi.
type AtoiResp struct {
	val int
	err error
}

// Val returns the value of the result.
func (r AtoiResp) Val() int {
	return r.val
}

// Err returns the error of the result.
func (r AtoiResp) Err() error {
	return r.err
}

// IsErr returns true if the result is an error.
func (r AtoiResp) IsErr() bool {
	return r.err != nil
}

// IsOk returns true if the result is ok.
func (r AtoiResp) IsOk() bool {
	return r.err == nil
}

// Atoi converts a string to an int.
func Atoi(s string) AtoiResp {
	val, err := strconv.Atoi(s)

	return AtoiResp{
		val: val,
		err: err,
	}
}
