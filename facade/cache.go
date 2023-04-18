package facade

import (
	"github.com/go-fires/framework/cache"
)

func Cache() *cache.Manager {
	var manager *cache.Manager
	if err := App().Make(cache.Cache, &manager); err != nil {
		panic(err)
	}

	return manager
}
