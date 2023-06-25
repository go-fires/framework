package tests

import (
	"github.com/go-fires/fires/tests/foundation"
	"github.com/go-fires/fires/tests/providers"
)

func App() *foundation.Application {
	app := foundation.NewApplication()

	app.Register(providers.NewHashingProvider(app))
	app.Register(providers.NewEncryptionProvider(app))

	return app
}
