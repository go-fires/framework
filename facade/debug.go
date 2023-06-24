package facade

import (
	"github.com/go-fires/fires/debug/recovery"
)

func DebugRecoveryHandler() recovery.Handler {
	return App().MustGet("debug.recovery.handler").(recovery.Handler)
}
