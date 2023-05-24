package socketio

import (
	"github.com/redis/go-redis/v9"
)

// RedisAdapterOptions is configuration to create new adapter
type RedisAdapterOptions struct {
	Prefix      string
	RedisClient redis.UniversalClient
}
