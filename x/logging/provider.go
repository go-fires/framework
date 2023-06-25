package logging

import (
	"github.com/go-fires/fires/x/contracts/container"
	foundation2 "github.com/go-fires/fires/x/contracts/foundation"
)

type Provider struct {
	app foundation2.Application
	*foundation2.UnimplementedProvider
}

func NewProvider(app foundation2.Application) *Provider {
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
