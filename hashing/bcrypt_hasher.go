package hashing

import (
	"golang.org/x/crypto/bcrypt"
)

var Bcrypt Hasher = &bcryptHasher{}

type bcryptHasher struct{}

// Make returns the hashed value.
func (b *bcryptHasher) Make(value string) (string, error) {
	return b.MakeWithCost(value, bcrypt.DefaultCost)
}

// MustMake returns the hashed value.
func (b *bcryptHasher) MustMake(value string) string {
	hashedValue, err := b.Make(value)

	if err != nil {
		panic(err)
	}

	return hashedValue
}

// MakeWithCost returns the hashed value with the given cost.
func (b *bcryptHasher) MakeWithCost(value string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), cost)

	if err != nil {
		return "", err
	}

	return string(bytes), err
}

// Check returns true if the value matches the hashed value.
func (b *bcryptHasher) Check(value, hashedValue string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(value)) == nil
}
