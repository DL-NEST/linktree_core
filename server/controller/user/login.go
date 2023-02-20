package user

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/service"
	"linktree_core/utils/result"
)

func (UserController) Login(ctx *gin.Context) {
	err, token := service.UserService.Login()
	if err != nil {
		return
	}
	result.OKWithResp(ctx, token)
}
