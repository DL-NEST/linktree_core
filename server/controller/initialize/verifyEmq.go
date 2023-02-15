package initialize

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/dto"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
)

// VerifyEmq 验证数据库连接
func (InitializeController) VerifyEmq(ctx *gin.Context) {
	var emq dto.Emq
	if err := ctx.ShouldBind(&emq); err != nil {
		result.APIResponse(ctx, code.ErrParam, err)
		return
	}
	//err := service.ServiceGroup.VerifyDBLink(emq)
	//if err != nil {
	//	result.APIResponse(ctx, code.ErrParam, "err")
	//	return
	//}
	result.APIResponse(ctx, code.OK, "")
}
