package logging

import (
	"github.com/go-fires/framework/contracts/container"
	"github.com/go-fires/framework/contracts/foundation"
)

type Provider struct {
	app foundation.Application
	*foundation.UnimplementedProvider
}

func NewProvider(app foundation.Application) *Provider {
	return &Provider{
		app: app,
	}
}

func (p *Provider) Register() {
	p.app.Singleton("logging", func(c container.Container) interface{} {
		if config, ok := p.app.Config().Get("logging").(*Config); ok {
			return NewManager(config)
		}

		panic("log config is not defined")
	})
}
