package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "linktree_core/docs"
	"linktree_core/server/middleware"
	"linktree_core/server/modules/result"
	"linktree_core/server/modules/result/code"
	"linktree_core/server/router"
	"net/http"
)

const rootPath = ""

// MainRouter 服务主路由
func MainRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	Router := gin.New()
	// 全局路由中间件
	Router.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.Cors(),
		middleware.Deprecation())
	// html
	Router.LoadHTMLGlob("./template/*.html")
	Router.StaticFS(rootPath+"/assets", http.Dir("./template/assets"))
	Router.StaticFile(rootPath+"/", "./template/index.html")
	// swagger生成的文档，更新 CMD swag init
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	RouterGroup := Router.Group(rootPath, middleware.JudgmentInit())
	// 不做鉴权的接口
	publicGroup := RouterGroup.Group("")
	{
		// 空路由
		Router.NoRoute(func(ctx *gin.Context) {
			result.APIResponse(ctx, code.ErrNotFound, "")
		})
		// 健康检查
		publicGroup.GET("/healthy", func(ctx *gin.Context) {
			result.APIResponse(ctx, code.OK, "")
		})
	}
	{
		router.BaseRouter.InitBaseRouter(publicGroup)             // 应用默认接口不用鉴权的登录等功能
		router.InitializeRouter.InitInitializeRouter(publicGroup) // 初始化应用的接口
	}
	// 做鉴权的接口
	privateGroup := RouterGroup.Group("", middleware.Auth())
	{
		router.ModelRouter.InitModelRouter(privateGroup)     // RESTAPI接口
		router.SysInfoRouter.InitSysInfoRouter(privateGroup) // 系统信息接口
	}

	return Router
}
