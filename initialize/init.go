package initialize

import (
	"fmt"
	"linktree_core/cmd/commands"
	"linktree_core/global"
	"linktree_core/modules/db"
	"linktree_core/modules/redis"
	"linktree_core/utils/pidFile"
	"os"
)

// InitApp 初始化程序打印版本号
func InitApp() {
	appInfo()
	// 读取配置文件
	readConfig()
	// 连接数据库和redis
	db.InitDBLink()
	redis.InitRedis()
}

// 输出应用信息检测版本更新
func appInfo() {
	var appVersion string
	// TODO 在这获取有没有版本更新,未发布待完善
	newVersion := "0.0.2"

	if commands.Version != newVersion {
		appVersion = appVersion + " \u001B[;34m-> " + newVersion + "\u001B[0m\u001B[0m\u001B[;32m " +
			"Github: https://github.com/DL-NEST/linktree_core \u001B[0m\n" +
			"\t\u001B[;34m检测到服务程序有新版本，需要更新程序请使用ctl命令：linktree updata\u001B[0m"
	} else {
		appVersion = appVersion + "\u001B[0m\u001B[0m\u001B[;32m " +
			"Github: https://github.com/DL-NEST/linktree_core \u001B[0m"
	}

	fmt.Printf("\u001B[;32m"+` 
                          _      _         _   _____                
                         | |    |_|       | | |_   _|               
                         | |     _  _ __  | | __| | _ __  ___   ___ 
                         | |    | || '_ \ | |/ /| || '__|/ _ \ / _ \
                         | |____| || | | ||   < | || |  |  __/|  __/
                         \_____/|_||_| |_||_|\_\\_/|_|   \___| \___|`+"\n\n\t"+
		"Version:\u001B[;35m %v \n", appVersion)
	fmt.Printf("\u001B[;32m" + "===========================================================================================\u001B[0m\n")
	global.GLOG.Infof("[ -conf %v, -log %v]", commands.ConfigPath, commands.LogPath)
	global.GLOG.Infof("pid:%v", os.Getpid())
	// 创建pid文件
	if _, err := pidFile.New("./pid"); err != nil {
		panic(err)
	}
}
