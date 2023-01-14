package control

import (
	"github.com/gin-gonic/gin"
	"linktree_server/server/result"
	"linktree_server/server/result/code"
	"linktree_server/server/service"
)

func InitUser(router *gin.RouterGroup) {
	router.GET("/user", GetUser)
	router.POST("/user", AddUser)
	router.PUT("/user", UpdateUser)
	router.DELETE("/user", DeleteUser)
}

func GetUser(ctx *gin.Context) {
	queryMap := ctx.Request.URL.Query()
	result.APIResponse(ctx, code.OK, service.UserService.GetUser(queryMap))
}

// AddUser 注册
func AddUser(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "POST")
}

// UpdateUser 注册
func UpdateUser(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "PUT")
}

// DeleteUser 注册
func DeleteUser(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "DELETE")
}
