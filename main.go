package main

import (
	"fmt"
	"github.com/spf13/viper"
	"linktree_core/bootstrap"
	"linktree_core/commands"
	"linktree_core/modules/amqp/emqx"
	"linktree_core/modules/database/db"
	"linktree_core/modules/database/redis"
	"linktree_core/server"
	"linktree_core/utils/glog"
	"linktree_core/utils/gos"
	"linktree_core/utils/pidfile"
	"os"
)

//go:generate swag init

func init() {
	commands.Execute()
}

func main() {
	switch commands.Mode {
	case "start":
		glog.Log.Debugf(os.Args[0])
		bootstrap.InitApp()
		// 读取配置文件
		if err := bootstrap.InitConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// 没有找到配置文件
				server.ConfigServe()
			} else {
				// 找到了但是出错了
				glog.Log.Errorf("读取配置文件失败:%v", err)
			}
		}
		// model 数据库
		db.CreateDBLink()
		redis.InitRedis()
		// 连接mq服务器
		emqx.LinkMqttBroker()
		bootstrap.OutInfo()
		// 启动web服务
		server.MainServe()
		break
	case "stop":
		err := pidfile.KillProcess()
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		fmt.Printf("Service stopped")
		break
	case "reboot":
		glog.Log.Debug("reboot")
		break
	case "update":
		glog.Log.Debug("update")
		break
	case "pwd":
		err, use, pwd := gos.ReadPassFile()
		if err != nil {
			fmt.Printf("There is no pass file")
			return
		}
		fmt.Printf("username:\t%s\n", use)
		fmt.Printf("password:\t%s\n", pwd)
		break
	}
}
