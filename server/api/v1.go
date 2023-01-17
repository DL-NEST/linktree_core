package api

import (
	"github.com/gin-gonic/gin"
	"linktree_core/server/control"
	"linktree_core/server/middleware"
)

func InjectV1(server *gin.Engine) {
	// V1的api
	V1 := server.Group("/v1", middleware.Auth())

	// user
	control.InitUser(V1)

	// sys /*系统状态的获取和操作
	sys := V1.Group("/sys")
	{
		sys.GET("/getSysInfo", control.GetSysInfo)
	}
	// file /*上传文件
	file := V1.Group("/file")
	{
		file.POST("/upload", control.UploadOne)
		file.POST("/uploadList", control.UploadList)
	}
	// device
	device := V1.Group("/device")
	device.GET("/getAllDevice", control.GetAllDevice)
	device.GET("/getAllDeviceGroup", control.GetAllDeviceGroup)
	device.GET("/getGroupDeviceList", control.GetGroupDeviceList)
	device.POST("/addDevice", control.AddDevice)

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
