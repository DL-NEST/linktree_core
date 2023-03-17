package initialize

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/model/dto"
	"linktree_core/server/modules/result/code"
	"linktree_core/server/service"
)

// CreateConfig
// @Tags      	Init
// @Summary   	创建配置文件
// @Description  创建配置文件
// @Security  	ApiKeyAuth
// @Accept    	application/json
// @Produce   	application/json
// @Param     	data  body    dto.SetupOpt	true	"config文件参数"
// @Success   	200  {object} result.Response
// @Failure   	404  {object} result.Response
// @Router    	/init/createConfig [post]
func (c InitializeController) CreateConfig(ctx *gin.Context) {
	var setupOpt dto.SetupOpt
	if err := c.New(ctx).BuildRequest(&setupOpt).Error(); err != nil {
		return
	}
	err := service.InitializeService.CreateConfig(setupOpt)
	if err != nil {
		c.Fail(code.ErrInit, err.Error())
		return
	}
	c.OK("")
}
