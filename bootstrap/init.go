package bootstrap

import (
	"fmt"
	"github.com/spf13/viper"
	"linktree_core/commands"
	"linktree_core/utils/glog"
	"linktree_core/utils/pidfile"

	"os"
)

// InitApp 初始化程序打印版本号
func InitApp() {
	var appVersion string
	// TODO 在这获取有没有版本更新
	newVersion := "0.0.2"

	if commands.Version != newVersion {
		appVersion = appVersion + " \u001B[;34m-> " + newVersion + "\u001B[0m\u001B[0m\u001B[;32m " +
			"Github: https://github.com/DL-NEST/linktree_server \u001B[0m\n" +
			"\t\u001B[;34m检测到服务程序有新版本，需要更新程序请使用命令：linktree updata\u001B[0m"
	} else {
		appVersion = appVersion + "\u001B[0m\u001B[0m\u001B[;32m " +
			"Github: https://github.com/DL-NEST/linktree_server \u001B[0m"
	}

	fmt.Printf("\u001B[;32m"+` 
                          _      _         _   _____                
                         | |    |_|       | | |_   _|               
                         | |     _  _ __  | | __| | _ __  ___   ___ 
                         | |    | || '_ \ | |/ /| || '__|/ _ \ / _ \
                         | |____| || | | ||   < | || |  |  __/|  __/
                         \_____/|_||_| |_||_|\_\\_/|_|   \___| \___|`+"\n\n\t"+
		"Version:\u001B[;35m %v \n", appVersion)
	fmt.Printf("\u001B[;32m" + "===========================================================================================\u001B[0m\n\n")
	glog.Log.Infof("[\u001B[;34m -conf \u001B[0m%v,\u001B[;34m -log \u001B[0m %v]", commands.ConfigPath, commands.LogPath)
	glog.Log.Infof("pid:%v", os.Getpid())
	// 创建pid文件
	if _, err := pidfile.New("./pid"); err != nil {
		panic(err)
	}
}

func OutInfo() string {
	port := viper.GetString("server.port")
	glog.Log.Info("监听服务端口:" + port)
	glog.Log.Debug("服务启动成功: http://localhost:" + port)
	glog.Log.Debug("swag文档地址: http://localhost:" + port + "/swagger/index.html\n")
	return port
}
