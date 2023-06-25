package logging

import (
	"fmt"
	logger2 "github.com/go-fires/fires/x/contracts/logger"
	"time"
)

type StdoutConfig struct {
}

type StdoutLogger struct {
	logger2.Loggerable
}

var _ logger2.Logger = (*StdoutLogger)(nil)

func NewStdoutLogger(name string) *StdoutLogger {
	return &StdoutLogger{
		Loggerable: func(level logger2.Level, message string) {
			fmt.Printf(
				"[%s] %s.%s: %s\n",
				time.Now().Format("2006-01-02 15:04:05"),
				name,
				level.UpperString(), message,
			)
		},
	}
}
