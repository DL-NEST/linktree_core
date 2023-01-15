package server

import (
	"github.com/gin-gonic/gin"
	"linktree_server/server/api"
	"linktree_server/server/middleware"
	"linktree_server/utils/result"
	"linktree_server/utils/result/code"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	// 全局路由
	server.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.Cors(),
		middleware.Deprecation())
	// 空路由
	server.NoRoute(func(c *gin.Context) {
		result.APIResponse(c, code.ErrNotFound, "")
	})
	// 固定路由
	api.Fixed(server)
	// v1路由
	api.InjectV1(server)

	return server
}
