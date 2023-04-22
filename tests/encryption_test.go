package tests

import (
	"testing"

	"github.com/go-fires/framework/facade"
	"github.com/stretchr/testify/assert"
)

func TestEncrypter(test *testing.T) {
	createApplication()

	ciphertext, err1 := facade.Encrypter().Encrypt("password")
	plaintext, err2 := facade.Encrypter().Decrypt(ciphertext)

	assert.Nil(test, err1)
	assert.Nil(test, err2)
	assert.True(test, plaintext == "password")
}
