package tests

import (
	"github.com/go-fires/framework/encryption"
	"github.com/go-fires/framework/facade"
	"github.com/go-fires/framework/foundation"
	"github.com/go-fires/framework/hashing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createApplication() *foundation.Application {
	return foundation.NewApplication()
}

func TestApp_Encrypter(test *testing.T) {
	app := createApplication()

	app.Register(encryption.NewProvider(app.Container))

	ciphertext, err1 := facade.Encrypter().Encrypt("password")
	plaintext, err2 := facade.Encrypter().Decrypt(ciphertext)

	assert.Nil(test, err1)
	assert.Nil(test, err2)
	assert.True(test, plaintext == "password")
}

func TestApp_Hasher(test *testing.T) {
	app := createApplication()

	app.Register(hashing.NewProvider(app.Container))

	assert.True(test, facade.Hash().Check("password", facade.Hash().MustMake("password")))
}
