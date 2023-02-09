package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"linktree_core/server/api"
	"linktree_core/server/control"
	"linktree_core/server/middleware"
	"linktree_core/utils/dlog"
	"linktree_core/utils/gos"
	"net/http"
)

// MainRouter 服务主路由
func mainRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// 全局路由
	router.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.Cors(),
		middleware.Deprecation())
	///csdfcv
	// 空路由
	router.NoRoute(control.NoRoute)
	// html
	router.LoadHTMLGlob("./template/*.html") // npm打包成dist的路径
	router.StaticFS("/assets", http.Dir("./template/assets"))
	router.StaticFile("/", "./template/index.html")
	// swagger生成的文档，更新 CMD swag init
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 固定路由
	api.Fixed(router)
	// v1路由
	api.InjectV1(router)
	// 测试
	api.Test(router)

	return router
}

// ConfigRouter 配置服务路由
func configRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// 全局路由
	router.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.Cors(),
		middleware.Deprecation()) // 空路由
	router.NoRoute(control.NoRoute)
	// 一些检测函数
	router.GET("/healthy", control.Healthy)
	// 登录函数
	router.POST("/login", control.FirstLogin)
	// 配置文件的接口
	cfg := router.Group("/config")
	control.InitCfg(cfg)

	return router
}

// MainServe 启动web服务
func MainServe() {
	router := mainRouter()
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
	router := configRouter()
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
