package facade

import "github.com/go-fires/framework/redis"

func Redis() *redis.Manager {
	var manager *redis.Manager
	if err := App().Make(redis.Redis, &manager); err != nil {
		panic(err)
	}

	return manager
}
