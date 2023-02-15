package main

import (
	"fmt"
	"linktree_core/bootstrap"
	"linktree_core/commands"
	"linktree_core/commands/flag"
	"linktree_core/global"
	"linktree_core/modules/db"
	"linktree_core/modules/redis"
	"linktree_core/server"
	"linktree_core/utils/gos"
	"os"
	"os/signal"
	"syscall"
)

//go:generate swag init

func init() {
	commands.Execute()
}

// @title                       linktree API
// @version                     0.1.0
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
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
	bootstrap.InitConfig()
	// 数据库
	db.InitDBLink()
	redis.InitRedis()

	// 连接mq服务器
	//emqx.LinkMqttBroker()
	// 连接kafka服务

	// 启动web服务
	server.StartServe()
	// 监听进行信号
	ListenerProcess()
}

func ListenerProcess() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	s := <-signalChan
	switch s {
	case syscall.SIGINT, syscall.SIGTERM:
		if o, ok := s.(syscall.Signal); ok {
			os.Exit(int(o))
		} else {
			os.Exit(0)
		}
	}
}

// backgroundStart 在后台运行服务
func backgroundStart() {
	fmt.Printf("Run the service in the background")
	gos.BackgroundStart()
}

// reboot 重启服务
func reboot() {
	global.GLOG.Debug("reboot")
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
	global.GLOG.Debug("update")
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
