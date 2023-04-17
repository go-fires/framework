package foundation

type Application interface {
	// Version Get the version number of the application.
	Version() string

	// Register a service provider with the application.
	Register(provider Provider)

	// Boot the application's service providers.
	Boot()

	// Terminate the application.
	Terminate()
}
