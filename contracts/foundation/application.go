package foundation

import "github.com/go-fires/framework/contracts/container"

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
}
