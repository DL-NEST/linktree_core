package control

import (
	"github.com/gin-gonic/gin"
	"linktree_server/server/result"
	"linktree_server/server/result/code"
)

func GetUser(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "db.QueryAllUser()")
}

// AddUser 注册
func AddUser(c *gin.Context) {

}

// UpdateUser 注册
func UpdateUser(c *gin.Context) {

}

// DeleteUser 注册
func DeleteUser(c *gin.Context) {

}
