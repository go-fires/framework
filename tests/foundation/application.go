package foundation

import (
	"github.com/go-fires/fires/encryption"
	"github.com/go-fires/fires/foundation"
	"github.com/go-fires/fires/hashing"
	"sync"
)

type Application struct {
	providers []foundation.Provider
	rw        sync.RWMutex

	Hasher    func() hashing.Hasher
	Encrypter func() *encryption.Encrypter
}

var _ foundation.Application = (*Application)(nil)

func NewApplication() *Application {
	return &Application{
		providers: make([]foundation.Provider, 0),
	}
}

func (app *Application) Register(provider foundation.Provider) {
	provider.Register()

	app.rw.Lock()
	defer app.rw.Unlock()
	app.providers = append(app.providers, provider)
}

func (app *Application) Boot() {
	app.rw.RLock()
	defer app.rw.RUnlock()

	for _, provider := range app.providers {
		provider.Boot()
	}
}

func (app *Application) Terminate() {
	app.rw.RLock()
	defer app.rw.RUnlock()

	for _, provider := range app.providers {
		provider.Terminate()
	}
}
