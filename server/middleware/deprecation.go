package middleware

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/modules/result"
	"linktree_core/server/modules/result/code"
)

var deprecationApi = []string{
	"/v1/*",
}

func Deprecation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, path := range deprecationApi {
			if ctx.Request.URL.Path == path {
				result.APIResponse(ctx, code.ErrOldAPI, "")
				ctx.Abort()
			}
		}
		ctx.Next()
	}
}
