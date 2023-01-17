package main

import (
	"github.com/spf13/viper"
	"linktree_core/bootstrap"
	"linktree_core/commands"
	"linktree_core/modules/db"
	"linktree_core/modules/emqx"
	"linktree_core/modules/redis"
	"linktree_core/server"
	"linktree_core/utils/logger"
)

//go:generate swag init

func init() {
	commands.Execute()
}

func main() {
	switch commands.Mode {
	case "start":
		// TODO: 判断是否第一次登录，跳转到初始化配置页面
		bootstrap.InitApp()
		bootstrap.InitConfig()
		// model
		db.CreateDBLink()
		redis.InitRedis()
		// mq
		emqx.LinkMqttBroker()
		// init over
		bootstrap.OutInfo()
		serverStart()
		break
	case "update":
		logger.Log.Debug("update")
		break
	case "reboot":
		logger.Log.Debug("reboot")
		break
	}
}

func serverStart() {
	service := server.InitRouter()
	port := viper.GetString("server.port")
	// 启动服务
	if !viper.InConfig("server.port") {
		return
	}

	err := service.Run(":" + port)
	if err != nil {
		logger.Log.Panicf("服务启动失败:%v", err)
		return
	}
}
