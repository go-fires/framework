package facade

import "github.com/go-fires/framework/encryption"

func Encrypter() *encryption.Encrypter {
	var encrypter *encryption.Encrypter

	if err := App().Make(encryption.EncrypterName, &encrypter); err != nil {
		panic(err)
	}

	return encrypter
}
