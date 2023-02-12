package system

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/service"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
)

type SysInfoApiGroup struct {
}

// GetSysInfo 获取系统信息
func (SysInfoApiGroup) GetSysInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.SysInfoService.GetSysInfo())
}
func (SysInfoApiGroup) GetMemInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.SysInfoService.GetMemInfo())
}
func (SysInfoApiGroup) GetDiskInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.SysInfoService.GetDiskInfo())
}
func (SysInfoApiGroup) GetHostInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.SysInfoService.GetHostInfo())
}
func (SysInfoApiGroup) GetNetInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.SysInfoService.GetNetInfo())
}

//func GetMqttList(c *gin.Context) {
//	result.APIResponse(c, code.OK, service.GetMqttList(emqx.MqttClient))
//}
