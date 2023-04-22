package facade

import (
	"github.com/go-fires/framework/contracts/logger"
	"github.com/go-fires/framework/logging"
)

func Logging() *logging.Manager {
	return App().MustGet("logging").(*logging.Manager)
}

func Log(names ...string) logger.Logger {
	return Logging().Channel(names...)
}
