package router

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/controller"
)

type defaultRouter struct{}

func (defaultRouter) InitBaseRouter(Router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	DefaultRouterGroup := Router.Group("base", handlers...)
	{
		DefaultRouterGroup.POST("login", controller.UserController.Login)
		DefaultRouterGroup.POST("test", controller.UserController.Test)
	}
}
