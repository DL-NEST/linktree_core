package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"linktree_core/global"
	"time"
)

func InitRedis() {
	if !global.State.SysInit {
		return
	}
	global.GLOG.Infof("连接redis")
	// 缓存用户的登录信息
	global.RdGroup.RdAuth = linkRdb(
		global.Config.GetString("persistence.redis.addr"),
		global.Config.GetString("persistence.redis.password"), // no password set
		0, // use default DB
	)
	global.RdGroup.MqMsg = linkRdb(
		global.Config.GetString("persistence.redis.addr"),
		global.Config.GetString("persistence.redis.password"), // no password set
		0, // use default DB
	)
}

func linkRdb(addr string, pwd string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd, // no password set
		DB:       db,  // use default DB
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		global.GLOG.Warnf("redis connect ping failed, err:%s", err)
		return nil
	}
	return client
}
