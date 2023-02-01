package api

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/control"
	"linktree_core/server/middleware"
)

func InjectV1(server *gin.Engine) {
	// V1的api
	V1 := server.Group("/v1", middleware.Auth())
	// sys 系统状态的获取和操作
	control.InitSys(V1)
	// file 上传文件
	control.InitFile(V1)
	/*实体接口*/
	// user
	control.InitUser(V1)
	// device
	control.InitDevice(V1)

	// websocket
	socket := V1.Group("/socket")
	socket.GET("/linkSocket", control.UpgradeSocket)
	// 广播到单个
	//socket.GET("/getTheOne", func(context *gin.Context) {
	//	sock.WsPool.WriteOne("3bd68f68-dfdc-400f-bf3b-f7840edc0385", wsPool.Text, "发送到单个")
	//})
	//// 广播到用户组
	//socket.GET("/getTheMore", func(context *gin.Context) {
	//	userList := []string{"3bd68f68-dfdc-400f-bf3b-f7840edc0385", ""}
	//	sock.WsPool.WriteMore(userList, wsPool.Text, "发送到用户组")
	//})
	//// 广播到全部
	//socket.GET("/broadcast", func(context *gin.Context) {
	//	sock.WsPool.Broadcast(wsPool.Json, service.GetCpuInfo())
	//})
	//// 获取已经连接的用户列表
	//socket.GET("/poolList", func(context *gin.Context) {
	//	result.APIResponse(context, code.OK, sock.WsPool.GetPoolList())
	//})
}
