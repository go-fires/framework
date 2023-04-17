package tests

import (
	"github.com/go-fires/framework/encryption"
	"github.com/go-fires/framework/facade"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypter(test *testing.T) {
	app := createApplication()

	app.Register(encryption.NewProvider(app.Container))

	ciphertext, err1 := facade.Encrypter().Encrypt("password")
	plaintext, err2 := facade.Encrypter().Decrypt(ciphertext)

	assert.Nil(test, err1)
	assert.Nil(test, err2)
	assert.True(test, plaintext == "password")
}
