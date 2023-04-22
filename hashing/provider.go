package hashing

import (
	"github.com/go-fires/framework/contracts/container"
	"github.com/go-fires/framework/contracts/foundation"
)

const Hash = "hash"

type Provider struct {
	app foundation.Application
	*foundation.UnimplementedProvider
}

var _ foundation.Provider = (*Provider)(nil)

func NewProvider(app foundation.Application) *Provider {
	return &Provider{
		app: app,
	}
}

func (h *Provider) Register() {
	h.app.Singleton(Hash, func(c container.Container) interface{} {
		return NewManagerWithContainer(c)
	})
}
