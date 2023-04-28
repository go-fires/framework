package logging

import (
	"fmt"
	"time"

	"github.com/go-fires/framework/contracts/logger"
)

type StdoutConfig struct {
}

type StdoutLogger struct {
	logger.Loggerable
}

var _ logger.Logger = (*StdoutLogger)(nil)

func NewStdoutLogger(name string) *StdoutLogger {
	return &StdoutLogger{
		Loggerable: func(level logger.Level, message string) {
			fmt.Printf(
				"[%s] %s.%s: %s\n",
				time.Now().Format("2006-01-02 15:04:05"),
				name,
				level.UpperString(), message,
			)
		},
	}
}
