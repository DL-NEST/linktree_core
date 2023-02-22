package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"linktree_core/global"
	"linktree_core/server/modules/result"
	"linktree_core/server/modules/result/code"
	"regexp"
)

// JudgmentInit Determine whether the application is initialized
func JudgmentInit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		compile, _ := regexp.Compile("/init/*")
		initPath := compile.MatchString(ctx.Request.URL.Path)

		// 判断是否初始化
		if !global.SysInit {
			// 鉴权初始化路由
			if initPath {
				ctx.Next()
				return
			}
			result.APIResponse(ctx, code.ErrNotInit, "服务器未初始化")
			ctx.Abort()
		} else {
			if initPath {
				result.APIResponse(ctx, code.ErrAuth, fmt.Sprintf("应用已初始化"))
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}
