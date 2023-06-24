package tests

import (
	"github.com/go-fires/fires/cache"
	"github.com/go-fires/fires/encryption"
	"github.com/go-fires/fires/hashing"
	"github.com/go-fires/fires/redis"
	foundation2 "github.com/go-fires/fires/x/foundation"
	rdb "github.com/redis/go-redis/v9"
)

type application struct {
	*foundation2.Application
}

func createApplication() *application {
	app := &application{
		Application: foundation2.NewApplication(),
	}

	app.configure()
	app.register()

	return app
}

func (app *application) register() {
	app.Register(redis.NewProvider(app))
	app.Register(cache.NewProvider(app))
	app.Register(encryption.NewProvider(app))
	app.Register(hashing.NewProvider(app))
}

func (app *application) configure() {
	app.Configure("app", &foundation2.Config{
		Name:     "test",
		Env:      "testing",
		Debug:    true,
		Timezone: "UTC",
		Locale:   "en",
		Key:      "base64:RUFGQlNQQVhEQ0lPR1JVVk5FUlFHWFBZR1BOS1lBVE0=",
	})
	app.Configure("redis", &redis.Config{
		Default: "default",
		Connections: map[string]redis.Configable{
			"default": &rdb.Options{
				Addr: "localhost:6379",
			},
		},
	})
	app.Configure("cache", &cache.Config{
		Default: "default",
		Stores: map[string]cache.StoreConfigable{
			"default": &cache.MemoryStoreConfig{},
			"redis": &cache.RedisStoreConfig{
				Connection: "default",
				Prefix:     "cache",
			},
		},
	})
	app.Configure("hashing", &hashing.Config{
		Driver: "bcrypt",
	})

}
