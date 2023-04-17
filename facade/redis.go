package facade

import "github.com/go-fires/framework/redis"

func Redis() *redis.Manager {
	var rdm *redis.Manager
	if err := App().Make(redis.Redis, &rdm); err != nil {
		panic(err)
	}

	return rdm
}
