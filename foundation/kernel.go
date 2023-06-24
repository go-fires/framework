package foundation

type Kernel interface {
	// Bootstrap the application.
	Bootstrap()

	// Handle the application.
	Handle()

	// Terminate the application.
	Terminate()
}
