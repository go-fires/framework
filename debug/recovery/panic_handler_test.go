package recovery

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanicHandler_Report(t *testing.T) {
	p := &PanicHandler{}

	assert.Panics(t, func() {
		p.Report("test")
	})
}
