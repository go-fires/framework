package hashing

import (
	"crypto/sha1"
	"fmt"
	"github.com/go-fires/framework/contracts/hashing"
)

type Sha1Hasher struct {
}

var _ hashing.Hasher = (*Sha1Hasher)(nil)

// NewSha1Hasher creates a new sha1 hasher instance.
func NewSha1Hasher() *Sha1Hasher {
	return &Sha1Hasher{}
}

// Make generates a new hashed value.
func (s *Sha1Hasher) Make(value string) (string, error) {
	hashedValue := sha1.New().Sum([]byte(value))

	return fmt.Sprintf("%x", hashedValue), nil
}

// MustMake generates a new hashed value.
func (s *Sha1Hasher) MustMake(value string) string {
	hashedValue, err := s.Make(value)

	if err != nil {
		panic(err)
	}

	return hashedValue
}

// Check checks the given value and hashed value.
func (s *Sha1Hasher) Check(value, hashedValue string) bool {
	hv, err := s.Make(value)

	if err != nil {
		return false
	}

	return hv == hashedValue
}
