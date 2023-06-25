package hashing

import (
	"crypto/sha1"
	"fmt"
)

var Sha1 Hasher = &sha1Hasher{}

type sha1Hasher struct{}

// Make generates a new hashed value.
func (h *sha1Hasher) Make(value string) (string, error) {
	hashedValue := sha1.New().Sum([]byte(value))

	return fmt.Sprintf("%x", hashedValue), nil
}

// MustMake generates a new hashed value.
func (h *sha1Hasher) MustMake(value string) string {
	hashedValue, err := h.Make(value)

	if err != nil {
		panic(err)
	}

	return hashedValue
}

// Check checks the given value and hashed value.
func (h *sha1Hasher) Check(value, hashedValue string) bool {
	hv, err := h.Make(value)

	if err != nil {
		return false
	}

	return hv == hashedValue
}
