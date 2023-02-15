package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"linktree_core/global"
)

func InitRedis() {
	if !global.SysInit {
		return
	}
	// 缓存用户的登录信息
	global.RdGroup.RdAuth = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("persistence.redis.addr"),
		Password: viper.GetString("persistence.redis.password"), // no password set
		DB:       0,                                             // use default DB
	})
}
