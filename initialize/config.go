package initialize

import (
	_ "embed"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"linktree_core/cmd/commands"
	"linktree_core/global"
	"linktree_core/utils/pidFile"
	"os"
)

// 把默认配置文件编译进二进制文件中
//
//go:embed config_default.yaml
var defaultConfig []byte

// create defaultConfig file
func creatDefaultConfig() {
	file, err := os.OpenFile("./config.yaml", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		global.GLOG.Errorf("默认配置文件生成失败：%s", err)
	}
	_, err = file.Write(defaultConfig)
	if err != nil {
		global.GLOG.Errorf("默认配置文件写入失败：%s", err)
		return
	}
	defer file.Close()
}

// InitConfig 初始化配置文件
func InitConfig() {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(commands.ConfigPath)
	// 服务的默认端口
	config.SetDefault("server.port", 5523)

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 没有找到配置文件
			creatDefaultConfig()
			pidFile.FirstPassFile()
			global.State.SysInit = false
			_ = config.ReadInConfig()
			return
		} else {
			// 找到了但是出错了
			global.GLOG.Errorf("读取配置文件失败:%v", err)
		}
	} else {
		global.State.SysInit = true
	}
	// 监听配置文件
	config.OnConfigChange(func(e fsnotify.Event) {
		err := config.ReadInConfig() // 读取配置文件
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	})
	config.WatchConfig()
	global.Config = config
}
