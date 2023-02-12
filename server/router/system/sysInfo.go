package system

import (
	"github.com/gin-gonic/gin"
	v1 "linktree_core/server/api/v1"
)

type SysInfoRouter struct{}

func (SysInfoRouter) InitSysInfoRouter(router *gin.RouterGroup) {
	var SysInfoApiGroup = v1.ApiGroup.SystemApiGroup.SysInfoApiGroup

	// 系统信息
	sysInfoRouter := router.Group("sysInfo")
	{
		sysInfoRouter.GET("all", SysInfoApiGroup.GetSysInfo)
		sysInfoRouter.GET("net", SysInfoApiGroup.GetNetInfo)
		sysInfoRouter.GET("host", SysInfoApiGroup.GetHostInfo)
		sysInfoRouter.GET("mem", SysInfoApiGroup.GetMemInfo)
		sysInfoRouter.GET("disk", SysInfoApiGroup.GetDiskInfo)
	}
}
