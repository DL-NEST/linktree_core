package system

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/service"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
)

type InitializeApiGroup struct {
}

// VerifyDB 验证数据库连接
func (InitializeApiGroup) VerifyDB(ctx *gin.Context) {
	var dsn service.Dsn
	if err := ctx.ShouldBind(&dsn); err != nil {
		result.APIResponse(ctx, code.ErrParam, err)
		return
	}
	err := service.VerifyDBLink(dsn)
	if err != nil {
		result.APIResponse(ctx, code.ErrParam, "err")
		return
	}
	result.APIResponse(ctx, code.OK, "")
}

// VerifyRedis 验证redis的连接
func (InitializeApiGroup) VerifyRedis(ctx *gin.Context) {
	var rOpt service.RedisOpt
	if err := ctx.ShouldBind(&rOpt); err != nil {
		result.APIResponse(ctx, code.ErrParam, err)
		return
	}
	err := service.VerifyRedisLink(rOpt)
	if err != nil {
		result.APIResponse(ctx, code.ErrParam, "err")
		return
	}
	result.APIResponse(ctx, code.OK, "")
}

// CreateConfig 验证数据库连接
func (InitializeApiGroup) CreateConfig(ctx *gin.Context) {
	var setupOpt service.SetupOpt
	if err := ctx.ShouldBind(&setupOpt); err != nil {
		result.APIResponse(ctx, code.ErrParam, err)
		return
	}
	err := service.CreateConfig(setupOpt)
	if err != nil {
		result.APIResponse(ctx, code.ErrInit, "")
		return
	}
	result.APIResponse(ctx, code.OK, "")
}
