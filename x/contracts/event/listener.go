package event

// Listener interface
type Listener interface {
	Handle(event Event)
}

// ListenerFunc listener function
type ListenerFunc func(event Event)

// Handle event
func (f ListenerFunc) Handle(event Event) {
	f(event)
}
