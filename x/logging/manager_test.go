package logging

import (
	"github.com/go-fires/fires/x/contracts/logger"
	"testing"
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
