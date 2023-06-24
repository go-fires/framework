package facade

import "github.com/go-fires/fires/encryption"

func Encrypter() *encryption.Encrypter {
	return App().MustGet(encryption.EncrypterName).(*encryption.Encrypter)
}
