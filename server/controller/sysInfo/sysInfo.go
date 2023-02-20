package sysInfo

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/service"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
)

// GetSysInfo
// @Tags      	SysInfo
// @Summary   	获取系统信息
// @Description  获取系统信息
// @Security  	ApiKeyAuth
// @Accept    	application/json
// @Produce   	application/json
// @Success   	200  {object} result.Response{data=dto.SysInfo}
// @Failure   	404  {object} result.Response
// @Router    	/sysInfo/all [get]
func (SysInfoController) GetSysInfo(c *gin.Context) {
	result.APIResponse(c, code.OK, service.SysInfoService.GetSysInfo())
}
