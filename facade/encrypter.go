package facade

import "github.com/go-fires/framework/encryption"

func Encrypter() *encryption.Encrypter {
	return App().MustGet(encryption.EncrypterName).(*encryption.Encrypter)
}
