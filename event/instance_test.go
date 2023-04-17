package event

import "testing"

func TestDispatch(t *testing.T) {
	event1, listener11, listener12 := newTestEvent("test"), &Test1Listener{}, &Test2Listener{}

	Listen("test", listener11)
	Listen("test", listener12)

	Dispatch(event1)
}
