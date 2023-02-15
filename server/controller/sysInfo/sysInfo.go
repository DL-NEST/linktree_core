package sysInfo

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/service"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
)

// GetSysInfo 获取系统信息
func (SysInfoController) GetSysInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.SysInfoService.GetSysInfo())
}
