package router

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/controller"
)

type modelRouter struct{}

func (modelRouter) InitModelRouter(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	var ModelControllerGroup = controller.ModelController

	ModelRouterGroup := router.Group("model", handlers...)

	userRouterGroup := ModelRouterGroup.Group("user")
	// user
	{
		userRouterGroup.GET("", ModelControllerGroup.UserController.GetUser)
		userRouterGroup.POST("", ModelControllerGroup.UserController.AddUser)
		userRouterGroup.PUT("", ModelControllerGroup.UserController.UpdateUser)
		userRouterGroup.DELETE("", ModelControllerGroup.UserController.DeleteUser)
	}
	// device
	{

	}
}
