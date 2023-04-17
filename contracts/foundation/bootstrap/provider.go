package bootstrap

import "github.com/go-fires/framework/foundation"

type Privider struct {
	*foundation.Application
}

func NewProvider(app *foundation.Application) *Privider {
	return &Privider{
		Application: app,
	}
}

func (p *Privider) Bootstrap() {
	p.Application.Boot()
}
