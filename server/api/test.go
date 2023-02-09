package api

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/control"
)

func Test(server *gin.Engine) {
	server.Any("/test/*uri", control.InitTest())
}
