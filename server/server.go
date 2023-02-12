package server

import (
	"github.com/spf13/viper"
	"linktree_core/server/router"
	"linktree_core/utils/dlog"
	"linktree_core/utils/gos"
	"net/http"
)

// MainServe 启动web服务
func MainServe() {
	router := router.MainRouter()
	// 启动服务
	if !viper.InConfig("server.port") {
		return
	}
	port := viper.GetString("server.port")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	err := srv.ListenAndServe()
	if err != nil {
		dlog.Log.Panicf("服务启动失败:%v", err)
		return
	}
}

// ConfigServe 启动一个初始化配置文件的服务
func ConfigServe() {
	dlog.Log.Warnf("未找到配置文件,启动初始化服务")
	// 生成随机密码
	gos.FirstPassFile()
	router := router.InitRouter()
	port := viper.GetString("server.port")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	err := srv.ListenAndServe()
	if err != nil {
		dlog.Log.Panicf("服务启动失败:%v", err)
		return
	}
}
