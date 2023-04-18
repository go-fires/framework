package cache

import "errors"

var (
	// ErrKeyNotFound is returned when the key is not found.
	ErrKeyNotFound = errors.New("key not found")

	// ErrUnknown is returned when the error is unknown.
	ErrUnknown = errors.New("unknown error")
)
