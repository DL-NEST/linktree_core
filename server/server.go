package server

import (
	"linktree_core/global"
	"net/http"
)

func StartServe() {
	Router := MainRouter()
	// 启动服务
	port := global.Config.GetString("server.port")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: Router,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			global.GLOG.Panicf("服务启动失败:%v", err)
		}
	}()
	outInfo()
}

func outInfo() string {
	port := global.Config.GetString("server.port")
	global.GLOG.Info("监听服务端口:" + port)
	global.GLOG.Debug("服务启动成功: http://localhost:" + port)
	global.GLOG.Debug("swag文档地址: http://localhost:" + port + "/swagger/index.html\n")
	return port
}
