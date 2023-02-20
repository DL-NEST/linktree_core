package initialize

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/dto"
	"linktree_core/server/service"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
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
func (InitializeController) CreateConfig(ctx *gin.Context) {
	var setupOpt dto.SetupOpt
	if err := ctx.ShouldBind(&setupOpt); err != nil {
		result.APIResponse(ctx, code.ErrParam, err)
		return
	}
	err := service.InitializeService.CreateConfig(setupOpt)
	if err != nil {
		result.APIResponse(ctx, code.ErrInit, "")
		return
	}
	result.APIResponse(ctx, code.OK, "")
}
