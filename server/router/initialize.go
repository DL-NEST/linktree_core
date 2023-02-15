package router

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/controller"
)

type initializeRouter struct{}

func (initializeRouter) InitInitializeRouter(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	var InitializeControllerGroup = controller.InitializeController

	// 系统信息
	InitializeRouterGroup := router.Group("init", handlers...)
	{
		InitializeRouterGroup.POST("/verifyDB", InitializeControllerGroup.VerifyDB)
		InitializeRouterGroup.POST("/verifyRedis", InitializeControllerGroup.VerifyRedis)
		InitializeRouterGroup.POST("/createConfig", InitializeControllerGroup.CreateConfig)
	}
}
