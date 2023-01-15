package control

import (
	"github.com/gin-gonic/gin"
	"linktree_server/server/service"
	"linktree_server/utils/result"
	"linktree_server/utils/result/code"
)

func InitUser(router *gin.RouterGroup) {
	router.GET("/user", GetUser)
	router.POST("/user", AddUser)
	router.PUT("/user", UpdateUser)
	router.DELETE("/user", DeleteUser)
}

func GetUser(ctx *gin.Context) {
	queryMap := ctx.Request.URL.Query()
	err, data := service.UserService.GetUser(queryMap)
	if err != nil {
		result.APIResponse(ctx, code.OK, data)
	}
	result.APIResponse(ctx, code.OK, data)
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
