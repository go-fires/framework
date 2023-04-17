package redis

import (
	"github.com/go-fires/framework/container"
)

const Redis = "redis"

type Provider struct {
	*container.Container
}

func NewProvider(c *container.Container) *Provider {
	return &Provider{
		Container: c,
	}
}

func (r *Provider) Register() {
	r.Singleton(Redis, func(c *container.Container) interface{} {
		return New(nil) // TODO: add config
	})
}
