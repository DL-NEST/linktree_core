package initialize

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/dto"
	"linktree_core/server/service"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
)

// CreateConfig 验证数据库连接
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
