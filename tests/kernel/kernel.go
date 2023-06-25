package kernel

import (
	f "github.com/go-fires/fires/foundation"
	"github.com/go-fires/fires/tests/foundation"
)

type Kernel struct {
	app *foundation.Application
}

var _ f.Kernel = (*Kernel)(nil)

func NewHttpKernel(app *foundation.Application) *Kernel {
	return &Kernel{
		app: app,
	}
}

func (kernel *Kernel) Bootstrap() {
	kernel.app.Boot()
}

func (kernel *Kernel) Handle() {
	kernel.Bootstrap()
	defer kernel.Terminate()
}

func (kernel *Kernel) Terminate() {
	kernel.app.Terminate()
}
