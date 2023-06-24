package main

import "testing"

func TestHttpKernel(t *testing.T) {
	app := NewApplication()
	app.Register(NewLogProvider(app))

	k := NewHttpKernel(app)
	k.Handle()
}
