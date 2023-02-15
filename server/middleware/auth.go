package middleware

import (
	"github.com/gin-gonic/gin"
	"linktree_core/global"
)

// Auth 登录鉴权
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//logger.Log.Debug("auth")
		//logger.Log.Debug(ctx.GetHeader("Authorization"))
		if ctx.IsWebsocket() {
			global.GLOG.Debug("is")
		} else {
			global.GLOG.Debug("cwsfd")
		}
		ctx.Next()
	}
}
