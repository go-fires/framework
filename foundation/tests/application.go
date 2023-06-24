package main

import (
	"github.com/go-fires/fires/foundation"
	"sync"
)

type Application struct {
	providers      []foundation.Provider
	providerRWLock sync.RWMutex

	Log *Log
}

var _ foundation.Application = (*Application)(nil)

func NewApplication() *Application {
	return &Application{
		providers: make([]foundation.Provider, 0),
	}
}

func (app *Application) Register(provider foundation.Provider) {
	provider.Register()

	app.providerRWLock.Lock()
	app.providers = append(app.providers, provider)
	app.providerRWLock.Unlock()
}

func (app *Application) Boot() {
	app.providerRWLock.RLock()
	defer app.providerRWLock.RUnlock()

	for _, provider := range app.providers {
		provider.Boot()
	}
}

func (app *Application) Terminate() {
	app.providerRWLock.RLock()
	defer app.providerRWLock.RUnlock()

	for _, provider := range app.providers {
		provider.Terminate()
	}
}
