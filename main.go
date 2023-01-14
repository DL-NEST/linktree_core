package main

import (
	"github.com/spf13/viper"
	"linktree_server/bootstrap"
	"linktree_server/commands"
	"linktree_server/modules/db"
	"linktree_server/modules/emqx"
	"linktree_server/server"
	"linktree_server/utils/logger"
)

//go:generate swag init

func init() {
	commands.Execute()
}

func main() {
	switch commands.Mode {
	case "start":
		bootstrap.InitApp()
		bootstrap.InitConfig()

		emqx.LinkMqttBroker()
		db.CreateDBLink()

		bootstrap.OutInfo()
		serverStart()
		break
	case "update":

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
