package facade

import (
	"github.com/go-fires/framework/debug/recovery"
)

func DebugRecoveryHandler() recovery.Handler {
	var handler recovery.Handler

	if err := App().Make("debug.recovery.handler", &handler); err != nil {
		panic(err)
	}

	return handler
}
