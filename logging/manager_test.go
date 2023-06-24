package logging

import (
	"testing"

	"github.com/go-fires/fires/contracts/logger"
)

func TestManager(t *testing.T) {
	m := NewManager(&Config{
		Default: "default",
		Channels: map[string]logger.Logger{
			"default": NewStdoutLogger("test"),
		},
	})

	m.Channel().Info("test") //nolint:errcheck
	m.Info("test")           //nolint:errcheck
}
