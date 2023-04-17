package encryption

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEncrypter(t *testing.T) {
	assert.Equal(t, &Encrypter{
		key: []byte("test"),
	}, NewEncrypter("test"))
}

func TestEncrypter(t *testing.T) {
	e := NewEncrypter("EAFBSPAXDCIOGRUVNERQGXPYGPNKYATM")

	ciphertext, _ := e.Encrypt("test")
	plaintext, _ := e.Decrypt(ciphertext)

	println(ciphertext, plaintext)

	assert.NotNil(t, ciphertext)
	assert.NotNil(t, plaintext)
	assert.Equal(t, "test", plaintext)
}

func TestEncrypter_Error(t *testing.T) {
	e := NewEncrypter("test")

	_, err1 := e.Encrypt("test")
	assert.Error(t, err1)

	_, err2 := e.Decrypt("test")
	assert.Error(t, err2)
}

func TestEncrypter_Decrypt_Error(t *testing.T) {
	e := NewEncrypter("EAFBSPAXDCIOGRUVNERQGXPYGPNKYATM")

	_, err1 := e.Decrypt("j9_mcZXKVlInk8bbpBqJOpmDp")
	assert.Error(t, err1)

	_, err2 := e.Decrypt("MQ==")
	assert.Error(t, err2)
}
