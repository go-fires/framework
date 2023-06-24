package tests

import (
	"github.com/go-fires/fires/facade"
	"github.com/go-fires/fires/hashing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashing(test *testing.T) {
	app := createApplication()

	app.Register(hashing.NewProvider(app))

	assert.True(test, facade.Hash().Check("password", facade.Hash().MustMake("password")))
}
