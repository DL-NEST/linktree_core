package initialize

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/dto"
	"linktree_core/server/service"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
)

type request struct {
	dto.Dsn
}

// VerifyDB
// @Tags      	Init
// @Summary   	验证数据库连接
// @Description  验证数据库连接
// @Security  	ApiKeyAuth
// @Accept    	application/json
// @Produce   	application/json
// @Param     	data  body    dto.Dsn	true	"DB连接参数"
// @Success   	200  {object} result.Response
// @Failure   	404  {object} result.Response
// @Router    	/init/verifyDB [post]
func (InitializeController) VerifyDB(ctx *gin.Context) {
	var dsn request
	if err := ctx.ShouldBind(&dsn); err != nil {
		result.APIResponse(ctx, code.ErrParam, err)
		return
	}
	err := service.InitializeService.VerifyDBLink(dsn.Dsn)
	if err != nil {
		result.APIResponse(ctx, code.ErrParam, "err")
		return
	}
	result.APIResponse(ctx, code.OK, "")
}
