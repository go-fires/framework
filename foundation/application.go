package foundation

type Application interface {
	// Register a service provider with the application.
	Register(provider Provider)

	// Boot the applications service providers.
	Boot()

	// Terminate the application.
	Terminate()
}
