package foundation

type Kernel interface {
	Bootstrap()
	Handle()
	Terminate()
}
