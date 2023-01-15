package control

import (
	"github.com/gin-gonic/gin"
	"linktree_server/utils/result"
	"linktree_server/utils/result/code"
)

// UserLoginParam 登录接收参数
type UserLoginParam struct {
	Username string `form:"username" json:"username" binding:"max=30"`
	Password string `form:"password" json:"password" binding:"max=40"`
}

// UserLoginRouterParam 登录返回参
type UserLoginRouterParam struct {
	UserName string `form:"username" json:"username" binding:"max=30"`
	Token    string `form:"token" json:"token" binding:"max=40"`
}

type UserRegisterParam struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Tel      int    `form:"tel" json:"tel" binding:"required"`
}

func UserLogin(ctx *gin.Context) {
	//var userLoginParam UserLoginParam
	//if err := ctx.ShouldBind(&userLoginParam); err != nil {
	//	result.APIResponse(ctx, code.ErrParam, err)
	//	return
	//}
	//boo, msg, user := service.Login(userLoginParam.Username, userLoginParam.Password)
	//if boo {
	//	result.APIResponse(ctx, msg, &UserLoginRouterParam{
	//		UserName: user.UserInfo.UserName,
	//		Token:    user.Token,
	//	})
	//} else {
	//	result.APIResponse(ctx, msg, "")
	//}
	result.APIResponse(ctx, code.OK, "")
}

// UserRegister 注册
func UserRegister(c *gin.Context) {
	var userRegisterParam UserRegisterParam
	if err := c.ShouldBind(&userRegisterParam); err != nil {
		result.APIResponse(c, code.ErrParam, err)
		return
	}
	//DB.AddUser(&DB.User{
	//	ID:         uuid.NewV4(),
	//	Name:       userRegisterParam.UserName,
	//	Tel:        34234,
	//	CreateTime: time.Now(),
	//})
	result.APIResponse(c, code.OK, "res")
}
