package redis

import (
	"github.com/go-redis/redis/v8"
)

var RdbAuth *redis.Client

func InitRedis() {
	// 缓存用户的登录信息
	RdbAuth = redis.NewClient(&redis.Options{
		Addr:     "localhost:6377",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
