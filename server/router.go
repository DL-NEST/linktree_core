package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"linktree_server/server/api"
	"linktree_server/server/middleware"
	"linktree_server/server/result"
	"linktree_server/server/result/code"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
		// 自定义头
		AllowHeaders:     []string{"Origin", "access_token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	})
}

func InitRouter() *gin.Engine {
	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	server.Use(middleware.Logger(), gin.Recovery(), Cors())
	// 空路由
	server.NoRoute(func(c *gin.Context) {
		result.APIResponse(c, code.ErrRouter, "")
	})
	// 固定路由
	api.Fixed(server)
	// v1路由
	api.InjectV1(server)

	return server
}
