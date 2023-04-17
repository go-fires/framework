package config

import (
	"github.com/go-fires/framework/container"
	"github.com/go-fires/framework/contracts/foundation"
)

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

func (p *Provider) Register() {
	p.Instance("config", NewConfig())
}
