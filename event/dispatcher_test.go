package event

import (
	"github.com/go-fires/framework/contracts/event"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testEvent struct {
	name   string
	Stoped bool
}

var _ event.Event = (*testEvent)(nil)
var _ event.StoppableEvent = (*testEvent)(nil)

func newTestEvent(name string) event.Event {
	return &testEvent{name: name}
}

func (t *testEvent) Name() string {
	return t.name
}

func (t *testEvent) IsPropagationStopped() bool {
	return t.Stoped
}

func (t *testEvent) SetStop() {
	t.Stoped = true
}

type Test1Listener struct {
	Val  string
	stop bool
}

var _ event.Listener = (*Test1Listener)(nil)

func (t *Test1Listener) Handle(event event.Event) {
	t.Val = "Test1 Done"

	if t.stop {
		event.(*testEvent).SetStop()
	}
}

func (t *Test1Listener) Stop() {
	t.stop = true
}

type Test2Listener struct {
	Val string
}

var _ event.Listener = (*Test2Listener)(nil)

func (t *Test2Listener) Handle(event event.Event) {
	t.Val = "Test2 Done"
}

func TestDispatcher_Dispatch(t *testing.T) {
	d := NewDispatcher()

	// event finish
	event1, listener11, listener12 := newTestEvent("test1"), &Test1Listener{}, &Test2Listener{}

	d.Listen("test1")
	d.Listen("test1", listener11, listener12)

	d.Dispatch(event1)

	assert.Equal(t, "Test1 Done", listener11.Val)
	assert.Equal(t, "Test2 Done", listener12.Val)

	// event stop
	event2, listener21, listener22 := newTestEvent("test2"), &Test1Listener{}, &Test2Listener{}

	d.Listen("test2", listener21)
	d.Listen("test2", listener22)

	listener21.Stop()

	d.Dispatch(event2)

	assert.Equal(t, "Test1 Done", listener21.Val)
	assert.Equal(t, "", listener22.Val)
}

func TestFlush(t *testing.T) {
	d := NewDispatcher()

	listener1, listener2 := &Test1Listener{}, &Test2Listener{}

	d.Listen("test1", listener1)
	d.Listen("test1", listener2)

	d.Listen("test2", listener1)
	d.Listen("test2", listener2)

	assert.Equal(t, 2, len(d.GetListeners("test1")))
	assert.Equal(t, 2, len(d.GetListeners("test2")))

	// Flush
	d.Flush("test1")

	assert.Equal(t, 0, len(d.GetListeners("test1")))
	assert.Equal(t, 2, len(d.GetListeners("test2")))

	// FlushAll
	d.Listen("test1", listener1)
	d.Listen("test1", listener2)

	assert.Equal(t, 2, len(d.GetListeners("test1")))
	assert.Equal(t, 2, len(d.GetListeners("test2")))

	d.FlushAll()

	assert.Equal(t, 0, len(d.GetListeners("test1")))
	assert.Equal(t, 0, len(d.GetListeners("test2")))
}

func TestDispatcher_ListenerFunc(t *testing.T) {
	d := NewDispatcher()

	event1 := newTestEvent("test1")

	d.Listen("test1", event.ListenerFunc(func(event event.Event) {
		assert.Equal(t, event1, event)
	}))
}
