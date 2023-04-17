package foundation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstance(t *testing.T) {
	app := NewApplication()

	SetInstance(app)

	assert.Same(t, app, GetInstance())
}
