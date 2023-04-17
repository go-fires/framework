package debug

type PanicHandler interface {
	// Report the given panic.
	Report(interface{})
}
