package cache

import (
	"github.com/go-fires/framework/contracts/cache"
	"github.com/go-fires/framework/contracts/container"
	"github.com/go-fires/framework/contracts/foundation"
)

const Cache = "cache"

type Provider struct {
	container.Container

	*foundation.UnimplementedProvider
}

var _ foundation.Provider = (*Provider)(nil)

func NewProvider(c container.Container) *Provider {
	return &Provider{
		Container: c,
	}
}

func (r *Provider) Register() {
	r.Singleton(Cache, func(c container.Container) interface{} {
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
