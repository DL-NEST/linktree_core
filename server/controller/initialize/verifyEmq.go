package initialize

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/model/dto"
)

// VerifyEmq 验证emq连接
func (c InitializeController) VerifyEmq(ctx *gin.Context) {
	var emq dto.Emq
	if err := c.New(ctx).BuildRequest(&emq).Error(); err != nil {
		return
	}
	//err := service.ServiceGroup.VerifyDBLink(emq)
	//if err != nil {
	//	result.APIResponse(ctx, code.ErrParam, "err")
	//	return
	//}
	c.OK("")
}
