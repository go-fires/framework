package providers

import (
	f "github.com/go-fires/fires/foundation"
	"github.com/go-fires/fires/hashing"
	"github.com/go-fires/fires/tests/foundation"
	"sync"
)

type HashingProvider struct {
	app *foundation.Application
	*f.UnimplementedProvider

	hasher hashing.Hasher
	once   sync.Once
}

func NewHashingProvider(app *foundation.Application) *HashingProvider {
	return &HashingProvider{
		app: app,
	}
}

func (p *HashingProvider) Register() {
	p.app.Hasher = func() hashing.Hasher {
		p.once.Do(func() {
			p.hasher = hashing.Md5
		})

		return p.hasher
	}
}
