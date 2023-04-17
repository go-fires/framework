package redis

import (
	"github.com/go-fires/framework/config"
	"github.com/go-fires/framework/container"
	"github.com/go-fires/framework/contracts/foundation"
	"github.com/redis/go-redis/v9"
)

const Redis = "redis"

type Provider struct {
	*container.Container
	*foundation.UnimplementedProvider
}

var _ foundation.Provider = (*Provider)(nil)

func NewProvider(c *container.Container) *Provider {
	return &Provider{
		Container: c,
	}
}

func (r *Provider) Register() {
	r.Singleton(Redis, func(c *container.Container) interface{} {
		if cfg, ok := r.getConfig().Get("redis").(*Config); ok {
			return New(cfg)
		}

		return New(r.defaultConfig())
	})
}

// config returns the framework configuration.
func (r *Provider) getConfig() *config.Config {
	var cfg *config.Config

	if err := r.Make("config", &cfg); err != nil {
		panic(err)
	}

	return cfg
}

// defaultConfig returns the default configuration for the redis provider.
func (r *Provider) defaultConfig() *Config {
	return &Config{
		Default: "default",
		Connections: map[string]Configable{
			"default": &redis.Options{
				Addr:     "localhost:6379",
				Password: "",
				DB:       0,
			},
		},
	}
}
