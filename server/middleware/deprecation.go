package middleware

import (
	"github.com/gin-gonic/gin"
	"linktree_core/global"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
)

var deprecationApi = []string{
	"/v1/:ss",
}

func Deprecation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		global.GLOG.Debug(ctx.GetHeader("Authorization"))
		for _, path := range deprecationApi {
			if ctx.Request.URL.Path == path {
				result.APIResponse(ctx, code.ErrOldAPI, "")
				ctx.Abort()
			}
		}
		ctx.Next()
	}
}
