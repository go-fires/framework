package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashing(t *testing.T) {
	app := App()

	assert.True(t, app.Hasher() == app.Hasher(), "Hasher should be a singleton")
	assert.True(t, app.Hasher().Check("test", app.Hasher().MustMake("test")))
}

func TestEncrypter(t *testing.T) {
	app := App()

	ciphertext, err := app.Encrypter().Encrypt("test")
	assert.Nil(t, err)

	plaintext, err := app.Encrypter().Decrypt(ciphertext)
	assert.Nil(t, err)

	assert.True(t, plaintext == "test")
}
