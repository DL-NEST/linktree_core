package middleware

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/modules/jwt"
	"linktree_core/server/modules/result"
	"linktree_core/server/modules/result/code"
)

// Auth 登录鉴权
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("Authorization")

		j := jwt.NewJWT()
		_, err := j.ParseToken(token)
		if err != nil {
			result.APIResponse(ctx, code.ErrAuth, "")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
