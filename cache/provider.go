package cache

import (
	"github.com/go-fires/framework/contracts/cache"
	"github.com/go-fires/framework/contracts/container"
	"github.com/go-fires/framework/contracts/foundation"
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
		return NewManager(&Config{
			Default: "redis",
			Stores: map[string]cache.StoreConfigable{
				"redis": &RedisStoreConfig{
					Connection: "cache",
					Prefix:     "cache",
				},
			},
		})
	})
}
