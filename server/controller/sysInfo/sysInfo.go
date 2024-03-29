package sysInfo

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/service"
)

// GetSysInfo
// @Tags      	SysInfo
// @Summary   	获取系统信息
// @Description  获取系统信息
// @Security  	ApiKeyAuth
// @Accept    	application/json
// @Produce   	application/json
// @Success   	200  {object} result.Response{data=dto.SysInfo} "返回系统信息"
// @Failure   	404  {object} result.Response
// @Router    	/sysInfo/all [get]
func (s SysInfoController) GetSysInfo(ctx *gin.Context) {
	s.New(ctx).OK(service.SysInfoService.GetSysInfo())
}
