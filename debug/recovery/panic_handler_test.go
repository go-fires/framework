package recovery

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanicHandler_Report(t *testing.T) {
	assert.Panics(t, func() {
		DefaultHandler.Report("test")
	})
}
