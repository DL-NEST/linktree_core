package api

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/control"
	"linktree_core/utils/logger"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
)

// Fixed 不需要验证就能的请求的api
func Fixed(server *gin.Engine) {

	// 一些检测函数
	server.GET("/healthy", func(ctx *gin.Context) {
		logger.Log.Debug(ctx.GetHeader("access_token"))
		result.APIResponse(ctx, code.OK, "ss")
	})
	// 公开,无Auth的接口
	server.POST("/login", control.UserLogin)
	server.POST("/register", control.UserRegister)

}
