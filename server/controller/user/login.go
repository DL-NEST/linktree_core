package user

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/service"
	"linktree_core/utils/result/code"
)

type request struct {
	UserName string `json:"userName,omitempty" binding:"required"`
	Password string `json:"password,omitempty"`
}

func (u UserController) Login(ctx *gin.Context) {
	var reo request
	err := u.New(ctx).BuildRequest(&reo).Error()
	if err != nil {
		return
	}
	err, token := service.UserService.Login()
	if err != nil {
		u.Fail(code.ErrInit, err.Error())
		return
	}
	u.OK(token)
}

func (u UserController) Test(ctx *gin.Context) {
	var req request
	err := u.New(ctx).BuildRequest(&req).Error()
	if err != nil {
		return
	}

	u.OK("成功")
}
