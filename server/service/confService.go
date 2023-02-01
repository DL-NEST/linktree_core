package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"linktree_core/utils/glog"
	"time"
)

type Dsn struct {
	Addr     string `json:"addr,omitempty"`
	Dataname string `json:"dataname,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type RedisOpt struct {
	Addr     string `json:"addr,omitempty"`
	Password string `json:"password,omitempty"`
}

type Emq struct {
	Addr     string `json:"addr,omitempty"`
	Client   string `json:"client,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type SetupOpt struct {
	Dsn      Dsn      `json:"dsn,omitempty"`
	RedisOpt RedisOpt `json:"redis_opt,omitempty"`
	Emq      Emq      `json:"emq,omitempty"`
}

func VerifyDBLink(ds Dsn) error {
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",
		ds.Username,
		ds.Password,
		ds.Addr,
		ds.Dataname,
		"utf8")
	// 打开数据库
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 1000,
	})
	if err != nil {
		return err
	}
	return nil
}

func VerifyRedisLink(rOpt RedisOpt) error {
	ctx := context.Background()
	c := redis.NewClient(&redis.Options{
		Addr:        rOpt.Addr,
		Password:    rOpt.Password,
		DB:          0,
		DialTimeout: 500 * time.Millisecond,
	})
	defer c.Close()
	_, err := c.Ping(ctx).Result()
	if err != nil {
		glog.Log.Debugf("%v", rOpt)
		return err
	}
	return nil
}

func CreateConfig(setupOpt SetupOpt) error {
	// 设置数据库-判断数据库连接类型
	if setupOpt.Dsn.Addr != "" {
		viper.SetDefault("persistence.mysql.addr", setupOpt.Dsn.Addr)
		viper.SetDefault("persistence.mysql.username", setupOpt.Dsn.Username)
		viper.SetDefault("persistence.mysql.password", setupOpt.Dsn.Password)
		viper.SetDefault("persistence.mysql.dataname", setupOpt.Dsn.Dataname)
		viper.SetDefault("persistence.mysql.charset", "utf8")
	} else {
		viper.SetDefault("persistence.mysql.addr", "")
		viper.SetDefault("persistence.mysql.username", "")
		viper.SetDefault("persistence.mysql.password", "")
		viper.SetDefault("persistence.mysql.dataname", "")
		viper.SetDefault("persistence.mysql.charset", "")
	}
	// 设置缓存方式-判断缓存类型
	if setupOpt.RedisOpt.Addr != "" {
		viper.SetDefault("persistence.redis.addr", setupOpt.RedisOpt.Addr)
		viper.SetDefault("persistence.redis.password", setupOpt.RedisOpt.Password)
	} else {
		viper.SetDefault("persistence.redis.addr", "")
	}
	// mqtt-broker
	viper.SetDefault("broker.addr", setupOpt.Emq.Addr)
	viper.SetDefault("broker.client", setupOpt.Emq.Client)
	viper.SetDefault("broker.username", setupOpt.Emq.Username)
	viper.SetDefault("broker.password", setupOpt.Emq.Password)

	err := viper.SafeWriteConfig()
	if err != nil {
		glog.Log.Errorf("创建配置文件失败:%v", err)
		return err
	}
	return nil
}
