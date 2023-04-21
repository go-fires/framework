package facade

import (
	"github.com/go-fires/framework/contracts/logger"
	"github.com/go-fires/framework/logging"
)

func Logging() *logging.Manager {
	var manager *logging.Manager

	if err := App().Make("logging", &manager); err != nil {
		panic(err)
	}

	return manager
}

func Log(names ...string) logger.Logger {
	return Logging().Channel(names...)
}
