package cache

import (
	"github.com/go-fires/fires/contracts/container"
	"github.com/go-fires/fires/contracts/foundation"
)

const Cache = "cache"

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

func (r *Provider) Register() {
	r.app.Singleton(Cache, func(c container.Container) interface{} {
		return NewManager(c)
	})
}
