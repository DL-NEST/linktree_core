package main

import (
	"fmt"
	"linktree_core/cmd/commands"
	"linktree_core/cmd/commands/flag"
	"linktree_core/global"
	"linktree_core/initialize"
	"linktree_core/modules/emqx"
	"linktree_core/server"
	"linktree_core/utils/daemon"
	"linktree_core/utils/pidFile"
	"linktree_core/utils/process"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	commands.Execute()
}

// @title                       linktree API
// @version                     0.1.0
// @description                 This is a sample Server pets
// @termsOfService  			https://github.com/DL-NEST/linktree_core
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:5523
// @BasePath  /
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
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
	initialize.InitApp()
	// 创建exhook-grpc-server
	emqx.CreateExHook()
	// 启动web服务
	server.StartServe()
	// 全局定时任务
	//worker.StartWorker()
	// 监听进程信号
	ListenerProcess()
}

func ListenerProcess() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	s := <-signalChan
	switch s {
	case syscall.SIGINT, syscall.SIGTERM:
		// 处理程序停止运行后的操作
		global.GLOG.Info("程序结束")
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
	daemon.BackgroundStart()
}

// reboot 重启服务
func reboot() {
	global.GLOG.Debug("reboot")
}

// stop 停止服务
func stop() {
	err := process.KillProcess()
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
	err, use, pwd := pidFile.ReadPassFile()
	if err != nil {
		fmt.Printf("There is no pass file")
		return
	}
	fmt.Printf("temporaryAccount \n")
	fmt.Printf("username:\t%s\n", use)
	fmt.Printf("password:\t%s\n", pwd)
}
