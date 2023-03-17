package initialize

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/model/dto"
	"linktree_core/server/modules/result/code"
	"linktree_core/server/service"
)

// VerifyRedis
// @Tags      	Init
// @Summary   	验证redis的连接
// @Description  验证redis的连接
// @Security  	ApiKeyAuth
// @Accept    	application/json
// @Produce   	application/json
// @Param     	data  body    dto.RedisOpt	true	"redis连接参数"
// @Success   	200  {object} result.Response
// @Failure   	404  {object} result.Response
// @Router    	/init/verifyRedis [post]
func (c InitializeController) VerifyRedis(ctx *gin.Context) {
	var rOpt dto.RedisOpt
	if err := c.New(ctx).BuildRequest(&rOpt).Error(); err != nil {
		return
	}
	err := service.InitializeService.VerifyRedisLink(rOpt)
	if err != nil {
		c.Fail(code.ErrParam, err.Error())
		return
	}
	c.OK("")
}
