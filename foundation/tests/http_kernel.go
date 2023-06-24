package main

import "github.com/go-fires/fires/foundation"

type HttpKernel struct {
	*Application
}

var _ foundation.Kernel = (*HttpKernel)(nil)

func NewHttpKernel(app *Application) *HttpKernel {
	return &HttpKernel{
		Application: app,
	}
}

func (kernel *HttpKernel) Bootstrap() {
	kernel.Application.Boot()
}

func (kernel *HttpKernel) Handle() {
	kernel.Bootstrap()
	defer kernel.Terminate()

	// Handle the test request.
	kernel.Log.Report("Handling the test request.")
}

func (kernel *HttpKernel) Terminate() {
	kernel.Application.Terminate()
}
