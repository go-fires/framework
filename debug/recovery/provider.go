package recovery

import (
	"github.com/go-fires/fires/contracts/container"
	"github.com/go-fires/fires/contracts/debug/recovery"
	"github.com/go-fires/fires/contracts/foundation"
)

type Provider struct {
	app foundation.Application
	*foundation.UnimplementedProvider

	handler recovery.Handler
}

type Option func(*Provider)

func NewProvider(app foundation.Application, opts ...Option) *Provider {
	p := &Provider{
		app: app,
	}

	p.handler = Handler{}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func WithHandler(handler recovery.Handler) Option {
	return func(p *Provider) {
		p.handler = handler
	}
}

func (p *Provider) Register() {
	p.app.Singleton("debug.recovery.handler", func(c container.Container) interface{} {
		return p.handler
	})
}
