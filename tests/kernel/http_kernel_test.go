package kernel

import (
	"github.com/go-fires/fires/tests/foundation"
	"testing"
)

func TestHttpKernel(t *testing.T) {
	app := foundation.NewApplication()

	k := NewHttpKernel(app)

	k.Handle()
}
