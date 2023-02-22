package result

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/modules/result/code"
)

// Response ...
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// APIResponse ...
func APIResponse(Ctx *gin.Context, start *code.Co, data any) {
	httpCo, co, message := code.Decode(start)

	if start == code.OK || start == code.OKorCreated {
		Ctx.JSON(httpCo, Response{
			Code: co,
			Msg:  message,
			Data: data,
		})
	} else {
		Ctx.JSON(httpCo, Response{
			Code: co,
			Msg:  message,
			Data: data,
		})
	}
}
