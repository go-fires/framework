package facade

import (
	cache2 "github.com/go-fires/fires/x/cache"
)

func Cache() *cache2.Manager {
	return App().MustGet(cache2.Cache).(*cache2.Manager)
}
