package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"linktree_core/server/middleware"
	"linktree_core/server/router/system"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
	"net/http"
)

const rootPath = ""

// MainRouter 服务主路由
func MainRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// 全局路由
	router.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.Cors(),
		middleware.Deprecation())

	// html
	router.LoadHTMLGlob("./template/*.html") // npm打包成dist的路径
	router.StaticFS(rootPath+"/assets", http.Dir("./template/assets"))
	router.StaticFile(rootPath+"/", "./template/index.html")
	// swagger生成的文档，更新 CMD swag init
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	systemRouter := system.SysRouterGroup

	// 不做鉴权的接口
	publicGroup := router.Group(rootPath)
	{
		// 空路由
		router.NoRoute(func(ctx *gin.Context) {
			result.APIResponse(ctx, code.ErrNotFound, "")
		})
		publicGroup.GET("/healthy", func(ctx *gin.Context) {
			result.APIResponse(ctx, code.OK, "")
		})
	}
	{
		systemRouter.InitModelRouter(publicGroup)
	}
	// 做鉴权的接口
	privateGroup := router.Group(rootPath)
	{
		systemRouter.InitSysInfoRouter(privateGroup)
	}

	return router
}

// InitRouter 配置服务路由
func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// 全局路由
	router.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.Cors(),
		middleware.Deprecation()) // 空路由

	systemRouter := system.SysRouterGroup

	// 不做鉴权的接口
	publicGroup := router.Group(rootPath)
	{
		// 空路由
		router.NoRoute(func(ctx *gin.Context) {
			result.APIResponse(ctx, code.ErrNotFound, "")
		})
		publicGroup.GET("/healthy", func(ctx *gin.Context) {
			result.APIResponse(ctx, code.OK, "")
		})
	}
	// 登录函数
	//router.POST("/login", api.FirstLogin)

	privateGroup := router.Group(rootPath)
	{
		systemRouter.InitSysInfoRouter(privateGroup)
	}

	return router
}
