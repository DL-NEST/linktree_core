package control

import (
	"github.com/gin-gonic/gin"
	"linktree_core/utils/result"
	"linktree_core/utils/result/code"
)

/*
* 设备的控制类
 */

func InitDevice(router *gin.RouterGroup) {
	router.GET("/device", GetDevice)
	router.POST("/device", AddDevice)
	router.PUT("/device", UpdateDevice)
	router.DELETE("/device", DeleteDevice)
}

// 1,设备的注册 2,设备的订阅主题设置 3,设备的控制 4,设备的查询

func GetDevice(ctx *gin.Context) {
	//queryMap := ctx.Request.URL.Query()
	//err, data := service.DeviceService.GetDevice(queryMap)
	//if err != nil {
	//	result.APIResponse(ctx, code.OK, data)
	//}
	//result.APIResponse(ctx, code.OK, data)
}

// AddDevice 注册
func AddDevice(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "POST")
}

// UpdateDevice 注册
func UpdateDevice(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "PUT")
}

// DeleteDevice 注册
func DeleteDevice(ctx *gin.Context) {
	result.APIResponse(ctx, code.OK, "DELETE")
}

// GetAllDevice 返回所有设备
func GetAllDevice(ctx *gin.Context) {
	//devices := db.QueryAllDevice()
	//result.APIResponse(ctx, code.OK, devices)
}

// GetAllDeviceGroup 返回所有设备组
func GetAllDeviceGroup(ctx *gin.Context) {
	//devices := db.QueryAllDeviceGroup()
	//logger.Log.Debug(devices)
	//result.APIResponse(ctx, code.OK, devices)
}

// GetGroupDeviceList 返回所有设备组
func GetGroupDeviceList(ctx *gin.Context) {
	//var groupName GroupName
	//err := ctx.ShouldBind(&groupName)
	//if err != nil {
	//	return
	//}
	//devices := db.QueryGroupDeviceList(groupName.GroupName)
	//result.APIResponse(ctx, code.OK, devices)
}

// AddDevice 添加设备
//func AddDevice(ctx *gin.Context) {
//	//var addDeviceParam AddDeviceParam
//	//if err := ctx.ShouldBind(&addDeviceParam); err != nil {
//	//	result.APIResponse(ctx, code.ErrParam, err)
//	//	return
//	//}
//	//// 表添加设备
//	//err1 := db.AddDevice(&db.Device{
//	//	DeviceID:    uuid.NewV4(),
//	//	DeviceName:  addDeviceParam.DeviceName,
//	//	DeviceGroup: addDeviceParam.DeviceGroup,
//	//	DeviceSubject: db.DeviceSubject{
//	//		DeviceState:   addDeviceParam.DeviceState,
//	//		DeviceControl: addDeviceParam.DeviceControl,
//	//	},
//	//	DeviceTemplate: addDeviceParam.DeviceSubjectName,
//	//})
//	//if err1 != nil {
//	//	result.APIResponse(ctx, code.ErrParam, err1)
//	//	return
//	//}
//	//result.APIResponse(ctx, code.OK, "ok")
//}
