package event

import (
	event2 "github.com/go-fires/fires/x/contracts/event"
	"sync"
)

// Dispatcher event dispatcher
type Dispatcher struct {
	listeners map[string][]event2.Listener

	mu sync.Mutex
}

var _ event2.Dispatcher = (*Dispatcher)(nil)

// NewDispatcher create new dispatcher
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		listeners: make(map[string][]event2.Listener, 0),
	}
}

// Listen add listener to event
func (d *Dispatcher) Listen(name string, listener ...event2.Listener) {
	if _, ok := d.listeners[name]; !ok {
		d.mu.Lock()
		defer d.mu.Unlock()

		d.listeners[name] = make([]event2.Listener, 0, len(listener))
	}

	d.listeners[name] = append(d.listeners[name], listener...)
}

// Dispatch event to all listeners
func (d *Dispatcher) Dispatch(e event2.Event) {
	for _, listener := range d.GetListeners(e.Name()) {
		if e, ok := e.(event2.StoppableEvent); ok && e.IsPropagationStopped() {
			return
		}

		listener.Handle(e)
	}
}

// GetListeners return all listeners of event
func (d *Dispatcher) GetListeners(name string) []event2.Listener {
	if listeners, ok := d.listeners[name]; ok {
		return listeners
	}

	return []event2.Listener{}
}

// Flush remove listeners of event
func (d *Dispatcher) Flush(name string) {
	delete(d.listeners, name)
}

// FlushAll remove all listeners
func (d *Dispatcher) FlushAll() {
	d.listeners = make(map[string][]event2.Listener, 0)
}
