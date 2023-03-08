package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type redisGroup struct {
	RdAuth *redis.Client
	MqMsg  *redis.Client
}

// 全局变量和句柄
var (
	// SysInit 应用是否初始化
	SysInit bool
	// DB 关系型数据库
	DB *gorm.DB
	// RdGroup 缓存型数据库
	RdGroup redisGroup
	// GLOG 全局日记输出
	GLOG *zap.SugaredLogger
	// GlogS 全局日记输出
	GlogS *zap.Logger
)
