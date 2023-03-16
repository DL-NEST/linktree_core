package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"linktree_core/global"
)

func InitRedis() {
	if !global.State.SysInit {
		return
	}
	global.GLOG.Infof("连接redis")
	// 缓存用户的登录信息
	global.RdGroup.RdAuth = linkRdb(
		viper.GetString("persistence.redis.addr"),
		viper.GetString("persistence.redis.password"), // no password set
		0, // use default DB
	)
	global.RdGroup.MqMsg = linkRdb(
		viper.GetString("persistence.redis.addr"),
		viper.GetString("persistence.redis.password"), // no password set
		1, // use default DB
	)
}

func linkRdb(addr string, pwd string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd, // no password set
		DB:       db,  // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GLOG.Warnf("redis connect ping failed, err:%s", err)
		return nil
	}
	return client
}
