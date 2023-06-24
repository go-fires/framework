package facade

import "github.com/go-fires/fires/config"

func Config() *config.Config {
	return App().MustGet("config").(*config.Config)
}
