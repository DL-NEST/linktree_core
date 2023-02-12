package model

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/service"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
)

type UserApiGroup struct {
}

func (UserApiGroup) GetUser(ctx *gin.Context) {
	queryMap := ctx.Request.URL.Query()
	err, data := service.UserService.GetUser(queryMap)
	if err != nil {
		result.APIResponse(ctx, code.OK, data)
	}
	result.APIResponse(ctx, code.OK, data)
}

// AddUser 注册
func (UserApiGroup) AddUser(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "POST")
}

// UpdateUser 注册
func (UserApiGroup) UpdateUser(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "PUT")
}

// DeleteUser 注册
func (UserApiGroup) DeleteUser(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "DELETE")
}
