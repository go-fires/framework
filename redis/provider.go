package redis

import (
	"github.com/go-fires/fires/contracts/container"
	"github.com/go-fires/fires/contracts/foundation"
)

const Redis = "redis"

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
	r.app.Singleton(Redis, func(c container.Container) interface{} {
		if config, ok := r.app.Config().Get("redis").(*Config); ok {
			return New(config)
		}

		return New(defaultConfig)
	})
}
