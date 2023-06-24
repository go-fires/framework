package facade

import "github.com/go-fires/fires/redis"

func Redis() *redis.Manager {
	return App().MustGet(redis.Redis).(*redis.Manager)
}
