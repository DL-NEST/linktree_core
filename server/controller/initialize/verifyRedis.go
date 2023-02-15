package initialize

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/dto"
	"linktree_core/server/service"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
)

// VerifyRedis 验证redis的连接
func (InitializeController) VerifyRedis(ctx *gin.Context) {
	var rOpt dto.RedisOpt
	if err := ctx.ShouldBind(&rOpt); err != nil {
		result.APIResponse(ctx, code.ErrParam, err)
		return
	}
	err := service.InitializeService.VerifyRedisLink(rOpt)
	if err != nil {
		result.APIResponse(ctx, code.ErrParam, "err")
		return
	}
	result.APIResponse(ctx, code.OK, "")
}
