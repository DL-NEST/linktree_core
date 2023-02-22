package model

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/modules/result"
	"linktree_core/server/modules/result/code"
)

type UserController struct {
}

// GetUser
// @Tags      	Model
// @Summary   	获取User
// @Description  获取User
// @Security  	ApiKeyAuth
// @Accept    	application/json
// @Produce   	application/json
// @Success   	200  {object} result.Response
// @Failure   	404  {object} result.Response
// @Router    	/model/user [get]
func (UserController) GetUser(ctx *gin.Context) {
	//queryMap := ctx.Request.URL.Query()
	//err, data := service.UserService.GetUser(queryMap)
	//if err != nil {
	//	result.APIResponse(ctx, code.OK, data)
	//}
	result.APIResponse(ctx, code.OK, "data")
}

// AddUser
// @Tags      	Model
// @Summary   	添加User
// @Description  添加User
// @Security  	ApiKeyAuth
// @Accept    	application/json
// @Produce   	application/json
// @Param     	data  body    dto.Dsn	true	"DB连接参数"
// @Success   	200  {object} result.Response
// @Failure   	404  {object} result.Response
// @Router    	/model/user [post]
func (UserController) AddUser(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "POST")
}

// UpdateUser
// @Tags      	Model
// @Summary   	更新User
// @Description  更新User
// @Security  	ApiKeyAuth
// @Accept    	application/json
// @Produce   	application/json
// @Param     	data  body    dto.Dsn	true	"DB连接参数"
// @Success   	200  {object} result.Response
// @Failure   	404  {object} result.Response
// @Router    	/model/user [put]
func (UserController) UpdateUser(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "PUT")
}

// DeleteUser
// @Tags      	Model
// @Summary   	删除User
// @Description  删除User
// @Security  	ApiKeyAuth
// @Accept    	application/json
// @Produce   	application/json
// @Param     	data  body    dto.Dsn	true	"DB连接参数"
// @Success   	200  {object} result.Response
// @Failure   	404  {object} result.Response
// @Router    	/model/user [delete]
func (UserController) DeleteUser(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "DELETE")
}
