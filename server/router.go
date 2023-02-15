package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "linktree_core/docs"
	"linktree_core/server/middleware"
	"linktree_core/server/router"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
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

	// 不做鉴权的接口
	publicGroup := Router.Group(rootPath, middleware.JudgmentInit())
	{
		// 空路由
		Router.NoRoute(func(ctx *gin.Context) {
			result.APIResponse(ctx, code.ErrNotFound, "")
		})
		publicGroup.GET("/healthy", func(ctx *gin.Context) {
			result.APIResponse(ctx, code.OK, "")
		})
	}
	{
		router.DefaultRouter.InitDefaultRouter(publicGroup)       // 应用默认接口
		router.InitializeRouter.InitInitializeRouter(publicGroup) // 初始化应用接口,本地密码鉴权
	}
	// 做鉴权的接口
	privateGroup := Router.Group(rootPath, middleware.JudgmentInit(), middleware.Auth())
	{
		router.ModelRouter.InitModelRouter(privateGroup)    // RESTAPI接口
		router.SysInfoRouter.InitSysInfoRouter(publicGroup) // 系统信息接口
	}

	return Router
}
