package logging

import "github.com/go-fires/framework/contracts/logger"

type Config struct {
	Default string

	Channels map[string]logger.Logger
}
