package event

type Dispatcher interface {
	Listen(name string, listener ...Listener)
	Dispatch(event Event)
	GetListeners(name string) []Listener
	Flush(name string)
	FlushAll()
}
