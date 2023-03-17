package user

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/model/dto/request"
	"linktree_core/server/modules/result/code"
	"linktree_core/server/service"
)

type LoginResponse struct {
	UserName string `json:"userName,omitempty"`
	UserId   uint   `json:"userId,omitempty"`
	Token    string `json:"token,omitempty"`
}

// Login
// @Tags      	Base
// @Summary   	用户登录
// @Description  用户登录
// @Accept    	application/json
// @Produce   	application/json
// @Param     	data  body    request.LoginRequest	true	"登录参数"
// @Success   	200  {object} result.Response{data=LoginResponse}
// @Failure   	404  {object} result.Response
// @Router    	/base/login [post]
func (c UserController) Login(ctx *gin.Context) {
	var reo request.LoginRequest
	if err := c.New(ctx).BuildRequest(&reo).Error(); err != nil {
		return
	}
	err, user, token := service.UserService.Login(reo)
	if err != nil {
		c.Fail(code.ErrLogin, err.Error())
		return
	}
	c.OK(LoginResponse{
		UserName: user.UserName,
		UserId:   user.ID,
		Token:    token,
	})
}
