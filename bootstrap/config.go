package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"linktree_core/commands"
)

// InitConfig 读取配置文件
func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(commands.ConfigPath)
	// 服务的默认端口
	viper.SetDefault("server.port", 5523)
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	// 监听配置文件
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	})
	viper.WatchConfig()
	return nil
}
