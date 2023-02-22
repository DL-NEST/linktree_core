package base

import (
	"github.com/gin-gonic/gin"
	"linktree_core/global"
	"linktree_core/server/modules/result"
	"linktree_core/server/modules/result/code"
)

type BaseController struct {
	Ctx         *gin.Context
	ErrorResult error
}

func (b *BaseController) New(ctx *gin.Context) *BaseController {
	b.Ctx = ctx
	return b
}

func (b *BaseController) BuildRequest(obj any) *BaseController {
	if b.Ctx == nil {
		global.GLOG.Warnf("The context is not bound")
		return b
	}
	err := b.Ctx.ShouldBind(obj)
	if err != nil {
		b.ErrorResult = err
		b.Fail(code.ErrParam, "")
	}
	return b
}

func (b *BaseController) BuildRequestUri(obj any) *BaseController {
	if b.Ctx == nil {
		global.GLOG.Warnf("The context is not bound")
		return b
	}
	err := b.Ctx.ShouldBindUri(obj)
	if err != nil {
		b.ErrorResult = err
		b.Fail(code.ErrParam, "")
	}
	return b
}

func (b *BaseController) Error() error {
	return b.ErrorResult
}

func (b *BaseController) OK(data any) {
	if b.Ctx == nil {
		global.GLOG.Warnf("The context is not bound")
		return
	}
	result.APIResponse(b.Ctx, code.OK, data)
}

func (b *BaseController) Fail(start *code.Co, data any) {
	if b.Ctx == nil {
		global.GLOG.Warnf("The context is not bound")
		return
	}
	result.APIResponse(b.Ctx, start, data)
}
