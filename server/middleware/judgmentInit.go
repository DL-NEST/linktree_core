package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"linktree_core/global"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
	"regexp"
)

// JudgmentInit Determine whether the application is initialized
func JudgmentInit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		compile, err := regexp.Compile("/init/*")
		if err != nil {
			return
		}
		// 判断是否初始化
		if !global.SysInit {
			// 鉴权初始化路由
			if compile.MatchString(ctx.Request.URL.Path) {

				ctx.Next()
				return
			}
			result.APIResponse(ctx, code.OK, fmt.Sprintf("%v", global.SysInit))
			ctx.Abort()
		}
		ctx.Next()
	}
}
