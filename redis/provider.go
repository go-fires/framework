package redis

import (
	"github.com/go-fires/framework/config"
	"github.com/go-fires/framework/contracts/container"
	"github.com/go-fires/framework/contracts/foundation"
	"github.com/redis/go-redis/v9"
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
		if cfg, ok := r.getConfig().Get("redis").(*Config); ok {
			return New(cfg)
		}

		return New(r.defaultConfig())
	})
}

// config returns the framework configuration.
func (r *Provider) getConfig() *config.Config {
	var cfg *config.Config

	if err := r.app.Make("config", &cfg); err != nil {
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
