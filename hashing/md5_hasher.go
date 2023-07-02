package hashing

import (
	"crypto/md5"
	"fmt"
)

var Md5 Hasher = &md5Hasher{}

type md5Hasher struct{}

// Make generates a new hashed value.
func (h *md5Hasher) Make(value string) (string, error) {
	hashedValue := md5.Sum([]byte(value))

	return fmt.Sprintf("%x", hashedValue), nil
}

// MustMake generates a new hashed value.
func (h *md5Hasher) MustMake(value string) string {
	hashedValue, err := h.Make(value)

	if err != nil {
		panic(err)
	}

	return hashedValue
}

// Check checks the given value and hashed value.
func (h *md5Hasher) Check(value, hashedValue string) bool {
	hv, err := h.Make(value)

	if err != nil {
		return false
	}

	return hv == hashedValue
}
