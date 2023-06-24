package event

import (
	"github.com/go-fires/fires/contracts/event"
	"sync"
)

// Dispatcher event dispatcher
type Dispatcher struct {
	listeners map[string][]event.Listener

	mu sync.Mutex
}

var _ event.Dispatcher = (*Dispatcher)(nil)

// NewDispatcher create new dispatcher
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		listeners: make(map[string][]event.Listener, 0),
	}
}

// Listen add listener to event
func (d *Dispatcher) Listen(name string, listener ...event.Listener) {
	if _, ok := d.listeners[name]; !ok {
		d.mu.Lock()
		defer d.mu.Unlock()

		d.listeners[name] = make([]event.Listener, 0, len(listener))
	}

	d.listeners[name] = append(d.listeners[name], listener...)
}

// Dispatch event to all listeners
func (d *Dispatcher) Dispatch(e event.Event) {
	for _, listener := range d.GetListeners(e.Name()) {
		if e, ok := e.(event.StoppableEvent); ok && e.IsPropagationStopped() {
			return
		}

		listener.Handle(e)
	}
}

// GetListeners return all listeners of event
func (d *Dispatcher) GetListeners(name string) []event.Listener {
	if listeners, ok := d.listeners[name]; ok {
		return listeners
	}

	return []event.Listener{}
}

// Flush remove listeners of event
func (d *Dispatcher) Flush(name string) {
	delete(d.listeners, name)
}

// FlushAll remove all listeners
func (d *Dispatcher) FlushAll() {
	d.listeners = make(map[string][]event.Listener, 0)
}
