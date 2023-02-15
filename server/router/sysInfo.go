package router

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/controller"
)

type sysInfoRouter struct{}

func (sysInfoRouter) InitSysInfoRouter(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	var SysInfoController = controller.SysInfoController

	// 系统信息
	SysInfoRouterGroup := router.Group("sysInfo")
	{
		SysInfoRouterGroup.GET("all", SysInfoController.GetSysInfo)
		//SysInfoRouter.GET("net", SysInfoApiGroup.GetNetInfo)
		//SysInfoRouter.GET("host", SysInfoApiGroup.GetHostInfo)
		//SysInfoRouter.GET("mem", SysInfoApiGroup.GetMemInfo)
		//SysInfoRouter.GET("disk", SysInfoApiGroup.GetDiskInfo)
	}
}
