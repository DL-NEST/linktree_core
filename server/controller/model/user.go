package model

import (
	"github.com/gin-gonic/gin"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
)

type UserController struct {
}

func (UserController) GetUser(ctx *gin.Context) {
	//queryMap := ctx.Request.URL.Query()
	//err, data := service.UserService.GetUser(queryMap)
	//if err != nil {
	//	result.APIResponse(ctx, code.OK, data)
	//}
	result.APIResponse(ctx, code.OK, "data")
}

// AddUser 注册
func (UserController) AddUser(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "POST")
}

// UpdateUser 注册
func (UserController) UpdateUser(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "PUT")
}

// DeleteUser 注册
func (UserController) DeleteUser(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "DELETE")
}
