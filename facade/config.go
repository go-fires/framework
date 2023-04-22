package facade

import "github.com/go-fires/framework/config"

func Config() *config.Config {
	return App().MustGet("config").(*config.Config)
}
