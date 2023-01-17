package control

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/service"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
)

// SysServer /* 获取系统信息

func GetSysInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.GetSysInfo())
}
func GetMemInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.GetMemInfo())
}
func GetDiskInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.GetDiskInfo())
}
func GetHostInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.GetHostInfo())
}
func GetNetInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.GetNetInfo())
}

//func GetMqttList(c *gin.Context) {
//	result.APIResponse(c, code.OK, service.GetMqttList(emqx.MqttClient))
//}
