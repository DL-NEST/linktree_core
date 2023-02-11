package main

import (
	"fmt"
	"github.com/spf13/viper"
	"linktree_core/bootstrap"
	"linktree_core/commands"
	"linktree_core/commands/flag"
	"linktree_core/modules/amqp/emqx"
	"linktree_core/modules/db"
	"linktree_core/modules/redis"
	"linktree_core/server"
	"linktree_core/utils/dlog"
	"linktree_core/utils/gos"
)

//go:generate swag init

func init() {
	commands.Execute()
}

func main() {
	switch commands.Mode {
	case flag.Start:
		appStart()
		break
	case flag.StartBg:
		backgroundStart()
		break
	case flag.Reboot:
		reboot()
		break
	case flag.Stop:
		stop()
		break
	case flag.Update:
		update()
		break
	case flag.Pwd:
		getPwd()
		break
	}
}

// appStart 服务启动
func appStart() {
	bootstrap.InitApp()
	// 读取配置文件
	if err := bootstrap.InitConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 没有找到配置文件
			server.ConfigServe()
		} else {
			// 找到了但是出错了
			dlog.Log.Errorf("读取配置文件失败:%v", err)
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
}

// backgroundStart 在后台运行服务
func backgroundStart() {
	fmt.Printf("Run the service in the background")
	gos.BackgroundStart()
}

// reboot 重启服务
func reboot() {
	dlog.Log.Debug("reboot")
}

// stop 停止服务
func stop() {
	err := gos.KillProcess()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("Service stopped")
}

// update 更新服务
func update() {
	dlog.Log.Debug("update")
}

// getPwd 获取初始密码
func getPwd() {
	err, use, pwd := gos.ReadPassFile()
	if err != nil {
		fmt.Printf("There is no pass file")
		return
	}
	fmt.Printf("username:\t%s\n", use)
	fmt.Printf("password:\t%s\n", pwd)
}
