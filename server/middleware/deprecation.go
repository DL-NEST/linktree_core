package middleware

import (
	"github.com/gin-gonic/gin"
	"linktree_server/utils/result"
	"linktree_server/utils/result/code"
)

var deprecationApi = []string{
	"/v1/:ss",
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
