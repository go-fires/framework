package event

import (
	"github.com/go-fires/framework/contracts/event"
	"sync"
)

var instance *Dispatcher
var once sync.Once

func GetDispatcher() *Dispatcher {
	if instance == nil {
		once.Do(func() {
			instance = NewDispatcher()
		})
	}

	return instance
}

func Listen(name string, listener event.Listener) {
	GetDispatcher().Listen(name, listener)
}

func Dispatch(event event.Event) {
	GetDispatcher().Dispatch(event)
}
