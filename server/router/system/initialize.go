package system

import (
	"github.com/gin-gonic/gin"
	v1 "linktree_core/server/api/v1"
)

type InitializeRouter struct{}

func (InitializeRouter) InitInitializeRouter(router *gin.RouterGroup) {
	var InitializeApiGroup = v1.ApiGroup.SystemApiGroup.InitializeApiGroup

	// 系统信息
	initializeRouter := router.Group("init")
	{
		initializeRouter.POST("/verifyDB", InitializeApiGroup.VerifyDB)
		initializeRouter.POST("/verifyRedis", InitializeApiGroup.VerifyRedis)
		initializeRouter.POST("/createConfig", InitializeApiGroup.CreateConfig)
	}
}
