package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局变量和句柄
var (
	// State 应用全局状态
	State state
	// Config 配置
	Config *viper.Viper
	// DB 关系型数据库
	DB *gorm.DB
	// RdGroup 缓存型数据库
	RdGroup redisGroup
	// GLOG 全局日记输出
	GLOG *zap.SugaredLogger
	// GlogS 全局日记输出
	GlogS *zap.Logger
)
