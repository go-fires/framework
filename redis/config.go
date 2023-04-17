package rediser

import "github.com/redis/go-redis/v9"

type Configable interface{}

var _ Configable = (*redis.Options)(nil)
var _ Configable = (*redis.ClusterOptions)(nil)
var _ Configable = (**redis.RingOptions)(nil)

type Config struct {
	Default string

	Connections map[string]Configable
}
