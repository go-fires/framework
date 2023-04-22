package facade

import (
	"github.com/go-fires/framework/hashing"
)

func Hash() *hashing.Manager {
	return App().MustGet(hashing.Hash).(*hashing.Manager)
}
