package recovery

type Handler interface {
	// Report the given panic.
	Report(interface{})
}
