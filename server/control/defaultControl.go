package control

import (
	"github.com/gin-gonic/gin"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
)

func Healthy(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "")
}

func NoRoute(ctx *gin.Context) {
	result.APIResponse(ctx, code.ErrNotFound, "")
}

func NoInit(ctx *gin.Context) {
	result.APIResponse(ctx, code.ErrNotFound, "The service is not initialized")
}
