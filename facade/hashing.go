package facade

import (
	"github.com/go-fires/framework/hashing"
)

func Hash() *hashing.Manager {
	var hash *hashing.Manager

	if err := App().Make(hashing.Hash, &hash); err != nil {
		panic(err)
	}

	return hash
}
