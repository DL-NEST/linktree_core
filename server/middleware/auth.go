package middleware

import (
	"github.com/gin-gonic/gin"
	"linktree_core/utils/dlog"
)

// Auth 登录鉴权
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//logger.Log.Debug("auth")
		//logger.Log.Debug(ctx.GetHeader("Authorization"))
		if ctx.IsWebsocket() {
			dlog.Log.Debug("is")
		} else {
			dlog.Log.Debug("cwsfd")
		}
		ctx.Next()
	}
}
