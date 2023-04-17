package tests

import (
	"github.com/go-fires/framework/facade"
	"github.com/go-fires/framework/hashing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashing(test *testing.T) {
	app := createApplication()

	app.Register(hashing.NewProvider(app.Container))

	assert.True(test, facade.Hash().Check("password", facade.Hash().MustMake("password")))
}
