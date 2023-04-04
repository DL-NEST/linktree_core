package global

import "github.com/go-redis/redis/v8"

type redisGroup struct {
	RdAuth *redis.Client
	MqMsg  *redis.Client
}
