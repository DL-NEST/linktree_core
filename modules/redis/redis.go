package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var RdbAuth *redis.Client

func InitRedis() {
	// 缓存用户的登录信息
	RdbAuth = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("persistence.redis.addr"),
		Password: viper.GetString("persistence.redis.password"), // no password set
		DB:       0,                                             // use default DB
	})
}
