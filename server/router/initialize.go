package router

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/controller"
	"linktree_core/server/middleware"
)

type initializeRouter struct{}

func (initializeRouter) InitInitializeRouter(router *gin.RouterGroup) {
	var InitializeControllerGroup = controller.InitializeController

	router.POST("init/login", InitializeControllerGroup.InitLogin)
	// 系统信息
	InitializeRouterGroup := router.Group("init", middleware.Auth())
	{
		InitializeRouterGroup.POST("/verifyDB", InitializeControllerGroup.VerifyDB)
		InitializeRouterGroup.POST("/verifyRedis", InitializeControllerGroup.VerifyRedis)
		InitializeRouterGroup.POST("/createConfig", InitializeControllerGroup.CreateConfig)
	}
}
