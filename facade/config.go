package facade

import "github.com/go-fires/framework/config"

func Config() *config.Config {
	var cfg *config.Config

	if err := App().Make("config", &cfg); err != nil {
		panic(err)
	}

	return cfg
}
