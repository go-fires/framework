package cache

import "github.com/go-fires/framework/contracts/cache"

type Config struct {
	Default string
	Stores  map[string]cache.StoreConfigable
}

var defaultConfig = &Config{
	Default: "default",
	Stores: map[string]cache.StoreConfigable{
		"default": &MemoryStoreConfig{},
		"redis": &RedisStoreConfig{
			Connection: "default",
			Prefix:     "cache",
		},
	},
}
