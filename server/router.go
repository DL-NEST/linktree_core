package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"linktree_core/server/api"
	"linktree_core/server/control"
	"linktree_core/server/middleware"
	"linktree_core/utils/glog"
	"linktree_core/utils/gos"
)

// MainRouter 服务主路由
func mainRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	// 全局路由
	server.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.Cors(),
		middleware.Deprecation())
	// 空路由
	server.NoRoute(control.NoRoute)
	// 固定路由
	api.Fixed(server)
	// v1路由
	api.InjectV1(server)

	return server
}

// ConfigRouter 配置服务路由
func configRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	// 全局路由
	server.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.Cors(),
		middleware.Deprecation())
	// 空路由
	server.NoRoute(control.NoRoute)
	// 一些检测函数
	server.GET("/healthy", control.Healthy)
	// 登录函数
	server.POST("/login", control.FirstLogin)
	// 配置文件的接口
	cfg := server.Group("/config")
	control.InitCfg(cfg)

	return server
}

// MainServe 启动web服务
func MainServe() {
	service := mainRouter()
	// 启动服务
	if !viper.InConfig("server.port") {
		return
	}
	port := viper.GetString("server.port")
	err := service.Run(":" + port)
	if err != nil {
		glog.Log.Panicf("服务启动失败:%v", err)
		return
	}
}

// ConfigServe 启动一个初始化配置文件的服务
func ConfigServe() {
	glog.Log.Warnf("未找到配置文件,启动初始化服务")
	// 生成随机密码
	gos.FirstPassFile()
	router := configRouter()
	err := router.Run(":" + viper.GetString("server.port"))
	if err != nil {
		glog.Log.Panicf("服务启动失败:%v", err)
		return
	}
}
