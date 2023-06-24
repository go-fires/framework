package facade

import (
	redis2 "github.com/go-fires/fires/x/redis"
)

func Redis() *redis2.Manager {
	return App().MustGet(redis2.Redis).(*redis2.Manager)
}
