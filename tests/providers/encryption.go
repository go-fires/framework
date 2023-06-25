package providers

import (
	"github.com/go-fires/fires/encryption"
	f "github.com/go-fires/fires/foundation"
	"github.com/go-fires/fires/tests/foundation"
	"sync"
)

type EncryptionProvider struct {
	app *foundation.Application
	*f.UnimplementedProvider

	encrypter *encryption.Encrypter
	once      sync.Once
}

func NewEncryptionProvider(app *foundation.Application) *EncryptionProvider {
	return &EncryptionProvider{
		app: app,
	}
}

func (p *EncryptionProvider) Register() {
	p.app.Encrypter = func() *encryption.Encrypter {
		p.once.Do(func() {
			p.encrypter = encryption.NewEncrypter("EAFBSPAXDCIOGRUVNERQGXPYGPNKYATM")
		})

		return p.encrypter
	}
}
