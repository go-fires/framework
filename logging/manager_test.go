package logging

import (
	"testing"

	"github.com/go-fires/framework/contracts/logger"
)

func TestManager(t *testing.T) {
	m := NewManager(&Config{
		Default: "default",
		Channels: map[string]logger.Logger{
			"default": NewStdoutLogger("test"),
		},
	})

	m.Channel().Info("test")
	m.Info("test")
}
