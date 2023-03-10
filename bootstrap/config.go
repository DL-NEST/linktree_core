package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"linktree_core/commands"
	"linktree_core/global"
	"linktree_core/utils/gos"
)

// InitConfig 读取配置文件
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(commands.ConfigPath)
	// 服务的默认端口
	viper.SetDefault("server.port", 5523)
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 没有找到配置文件
			gos.FirstPassFile()
			global.SysInit = false
			return
		} else {
			// 找到了但是出错了
			global.GLOG.Errorf("读取配置文件失败:%v", err)
		}
	} else {
		global.SysInit = true
	}
	// 监听配置文件
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	})
	viper.WatchConfig()
}
