package middleware

import (
	"github.com/gin-gonic/gin"
	"linktree_core/utils/glog"
)

// Auth 登录鉴权
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//logger.Log.Debug("auth")
		//logger.Log.Debug(ctx.GetHeader("Authorization"))
		if ctx.IsWebsocket() {
			glog.Log.Debug("is")
		} else {
			glog.Log.Debug("cwsfd")
		}
		ctx.Next()
	}
}
