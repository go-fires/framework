package redis

import (
	"github.com/go-fires/fires/x/contracts/container"
	foundation2 "github.com/go-fires/fires/x/contracts/foundation"
)

const Redis = "redis"

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
	r.app.Singleton(Redis, func(c container.Container) interface{} {
		if config, ok := r.app.Config().Get("redis").(*Config); ok {
			return New(config)
		}

		return New(defaultConfig)
	})
}
