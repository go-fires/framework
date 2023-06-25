package logging

import (
	"github.com/go-fires/fires/x/contracts/logger"
)

type Config struct {
	Default string

	Channels map[string]logger.Logger
}
