package initialize

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/model/dto"
	"linktree_core/server/modules/result/code"
	"linktree_core/server/service"
)

// VerifyDB
// @Tags      	Init
// @Summary   	验证数据库连接
// @Description  验证数据库连接
// @Security  	ApiKeyAuth
// @Accept    	application/json
// @Produce   	application/json
// @Param     	data  body    dto.Dsn	true	"DB连接参数"
// @Success   	200  {object} result.Response
// @Failure   	404  {object} result.Response
// @Router    	/init/verifyDB [post]
func (c InitializeController) VerifyDB(ctx *gin.Context) {
	var dsn dto.Dsn
	if err := c.New(ctx).BuildRequest(&dsn).Error(); err != nil {
		return
	}
	err := service.InitializeService.VerifyDBLink(dsn)
	if err != nil {
		c.Fail(code.ErrParam, err.Error())
		return
	}
	c.OK("")
}
