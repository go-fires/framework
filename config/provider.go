package config

import (
	"github.com/go-fires/framework/contracts/foundation"
)

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

func (p *Provider) Register() {
	p.app.Instance("config", NewConfig())
}
