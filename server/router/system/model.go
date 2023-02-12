package system

import (
	"github.com/gin-gonic/gin"
	v1 "linktree_core/server/api/v1"
)

type ModelRouter struct{}

func (ModelRouter) InitModelRouter(router *gin.RouterGroup) {
	var ModelApiGroup = v1.ApiGroup.ModelApiGroup

	modelRouter := router.Group("model")

	userRouter := modelRouter.Group("user")
	// user
	{
		userRouter.GET("", ModelApiGroup.UserApiGroup.GetUser)
		userRouter.POST("", ModelApiGroup.UserApiGroup.AddUser)
		userRouter.PUT("", ModelApiGroup.UserApiGroup.UpdateUser)
		userRouter.DELETE("", ModelApiGroup.UserApiGroup.DeleteUser)
	}
	// device
	{

	}
}
