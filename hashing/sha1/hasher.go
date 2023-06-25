package sha1

import (
	"crypto/sha1"
	"fmt"
	"github.com/go-fires/fires/hashing"
)

var global = New()

func New() hashing.Hasher {
	return &hasher{}
}

func Make(value string) (string, error) {
	return global.Make(value)
}

func MustMake(value string) string {
	return global.MustMake(value)
}

func Check(value, hashedValue string) bool {
	return global.Check(value, hashedValue)
}

type hasher struct{}

var _ hashing.Hasher = (*hasher)(nil)

// Make generates a new hashed value.
func (h *hasher) Make(value string) (string, error) {
	hashedValue := sha1.New().Sum([]byte(value))

	return fmt.Sprintf("%x", hashedValue), nil
}

// MustMake generates a new hashed value.
func (h *hasher) MustMake(value string) string {
	hashedValue, err := h.Make(value)

	if err != nil {
		panic(err)
	}

	return hashedValue
}

// Check checks the given value and hashed value.
func (h *hasher) Check(value, hashedValue string) bool {
	hv, err := h.Make(value)

	if err != nil {
		return false
	}

	return hv == hashedValue
}
