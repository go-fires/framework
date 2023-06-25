package foundation

import (
	"github.com/go-fires/fires/config"
	"github.com/go-fires/fires/x/contracts/container"
)

type Application interface {
	container.Container

	// Version Get the version number of the application.
	Version() string

	// Register a service provider with the application.
	Register(provider Provider)

	// Boot the application's service providers.
	Boot()

	// Terminate the application.
	Terminate()

	// Configure the real-time facade namespace.
	Configure(name string, value interface{})

	// Config Get the configuration repository instance.
	Config() *config.Config
}
