package server

import (
	"github.com/spf13/viper"
	"linktree_core/global"
	"net/http"
)

func StartServe() {
	Router := MainRouter()
	// 启动服务
	if !viper.InConfig("server.port") {
		return
	}
	port := viper.GetString("server.port")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: Router,
	}
	err := srv.ListenAndServe()
	if err != nil {
		global.GLOG.Panicf("服务启动失败:%v", err)
		return
	}
}
