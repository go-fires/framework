package cache

import (
	"github.com/go-fires/fires/x/contracts/container"
	foundation2 "github.com/go-fires/fires/x/contracts/foundation"
)

const Cache = "cache"

type Provider struct {
	app foundation2.Application

	*foundation2.UnimplementedProvider
}

var _ foundation2.Provider = (*Provider)(nil)

func NewProvider(app foundation2.Application) *Provider {
	return &Provider{
		app: app,
	}
}

func (r *Provider) Register() {
	r.app.Singleton(Cache, func(c container.Container) interface{} {
		return NewManager(c)
	})
}
