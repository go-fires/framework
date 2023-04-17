package foundation

import (
	"github.com/go-fires/framework/contracts/debug"
	"github.com/go-fires/framework/contracts/foundation"
	"github.com/go-fires/framework/foundation/bootstrap"
)

var (
	bootstrappers = []foundation.Bootstrapper{
		&bootstrap.BootPrivider{},
	}
)

type Kernel struct {
	app foundation.Application
}

var _ foundation.Kernel = (*Kernel)(nil)

func NewKernel(app foundation.Application) *Kernel {
	return &Kernel{
		app: app,
	}
}

func (k *Kernel) Bootstrap() {
	for _, bootstrapper := range bootstrappers {
		bootstrapper.Bootstrap(k.app)
	}
}

func (k *Kernel) Handle() {
	defer func() {
		if err := recover(); err != nil {
			k.reportPanic(err.(error))
		}
	}()

	k.Bootstrap()
}

func (k *Kernel) Terminate() {
	k.app.Terminate()
}

func (k *Kernel) reportPanic(p interface{}) {
	var handler debug.PanicHandler

	if err := k.app.Make("panic", &handler); err != nil {
		return
	}

	handler.Report(p)
}
