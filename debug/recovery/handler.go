package recovery

type Handler interface {
	// The Report the given panic.
	Report(interface{})
	// ShouldReport returns true if the given panic should be reported.
	ShouldReport(interface{}) bool
}
