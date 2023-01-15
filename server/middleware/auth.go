package middleware

import (
	"github.com/gin-gonic/gin"
)

// Auth 登录鉴权
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//logger.Log.Debug("auth")
		//logger.Log.Debug(ctx.GetHeader("Authorization"))
		ctx.Next()
	}
}
