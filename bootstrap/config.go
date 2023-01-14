package bootstrap

import (
	"github.com/spf13/viper"
	"linktree_server/commands"
	"linktree_server/utils/logger"
)

// InitConfig 读取配置文件
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(commands.ConfigPath)
	// 设置默认值
	defaultConfig()
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Log.Warn("未指定配置文件,在目录下生成配置文件")
			// 没有配置文件的话，会自动生成一个配置文件
			err1 := viper.SafeWriteConfig()
			if err1 != nil {
				return
			}
		} else {
			// 找到了但是出错了
			logger.Log.Errorf("读取配置文件失败:%v", err)
		}
	}
}

// defaultConfig 配置文件的默认配置
func defaultConfig() {
	viper.SetDefault("persistence.redis.addr", "127.0.0.1:6379")
	viper.SetDefault("persistence.mysql.addr", "127.0.0.1:3306")
	viper.SetDefault("persistence.mysql.username", "")
	viper.SetDefault("persistence.mysql.password", "")
	viper.SetDefault("persistence.mysql.dataname", "")
	viper.SetDefault("persistence.mysql.charset", "")
	viper.SetDefault("broker.client", "go-server")
	viper.SetDefault("broker.password", "")
	viper.SetDefault("broker.username", "")
	viper.SetDefault("broker.addr", "127.0.0.1:1883")
	viper.SetDefault("server.port", "5523")
	viper.SetDefault("server.host", "127.0.0.1")
}
