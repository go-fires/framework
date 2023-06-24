package facade

import (
	"github.com/go-fires/fires/cache"
)

func Cache() *cache.Manager {
	return App().MustGet(cache.Cache).(*cache.Manager)
}
