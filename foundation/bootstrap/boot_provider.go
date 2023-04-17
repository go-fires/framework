package bootstrap

import "github.com/go-fires/framework/contracts/foundation"

type BootPrivider struct {
}

var _ foundation.Bootstrapper = (*BootPrivider)(nil)

func (p *BootPrivider) Bootstrap(app foundation.Application) {
	app.Boot()
}
