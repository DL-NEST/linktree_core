package api

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/control"
)

// Fixed 不需要验证就能的请求的api
func Fixed(server *gin.Engine) {

	// 一些检测函数
	server.GET("/healthy", control.Healthy)
	// 公开,无Auth的接口
	server.POST("/login", control.UserLogin)
	server.POST("/register", control.UserRegister)

	// 测试接口

}
