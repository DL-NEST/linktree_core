package router

import (
	"github.com/gin-gonic/gin"
)

type defaultRouter struct{}

func (defaultRouter) InitDefaultRouter(Router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	//defaultApi := v1.ApiGroup.SystemApiGroup
	//DefaultRouter := Router.Group("base")
	//{
	//	DefaultRouter.POST("login", defaultApi.GetDiskInfo)
	//}
}
