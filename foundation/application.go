package foundation

import (
	"github.com/go-fires/framework/container"
	"github.com/go-fires/framework/contracts/foundation"
	"sync"
)

const Version = "0.0.1"

type Application struct {
	*container.Container

	providers []foundation.Provider
	mu        sync.RWMutex
	booted    bool
}

var _ foundation.Application = (*Application)(nil)

func NewApplication() *Application {
	app := &Application{
		Container: container.NewContainer(),
		providers: make([]foundation.Provider, 10),
	}

	app.init()

	return app
}

func (a *Application) init() {
	a.Container.Instance("app", a)

	SetInstance(a)
}

func (a *Application) Version() string {
	return Version
}

func (a *Application) Register(provider foundation.Provider) {
	a.mu.Lock()
	defer a.mu.Unlock()

	provider.Register()

	a.providers = append(a.providers, provider)

	if a.booted {
		a.bootProvider(provider)
	}
}

func (a *Application) Terminate() {
	if !a.booted {
		return
	}

	for _, p := range a.providers {
		a.terminateProvider(p)
	}
}

func (a *Application) terminateProvider(p foundation.Provider) {
	if p == nil {
		return
	}

	p.Terminate()
}

func (a *Application) Boot() {
	if a.booted {
		return
	}

	for _, p := range a.providers {
		a.bootProvider(p)
	}

	a.booted = true
}

func (a *Application) bootProvider(provider foundation.Provider) {
	if provider == nil {
		return
	}

	provider.Boot()
}
