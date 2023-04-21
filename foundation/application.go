package foundation

import (
	"sync"

	"github.com/go-fires/framework/config"
	"github.com/go-fires/framework/container"
	"github.com/go-fires/framework/contracts/foundation"
	"github.com/go-fires/framework/foundation/providers"
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

	a.registerBaseProviders()
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

func (a *Application) registerBaseProviders() {
	a.Register(providers.NewConfigProvider(a))
}

func (a *Application) Configure(name string, value interface{}) {
	a.Config().Set(name, value)
}

func (a *Application) Config() *config.Config {
	var cfg *config.Config

	if a.Make("config", &cfg) != nil {
		panic("config not found")
	}

	return cfg
}
