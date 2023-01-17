package control

import (
	"github.com/gin-gonic/gin"
	"linktree_core/modules/db/model"
)

/*
* 设备的控制类
 */

// AddDeviceParam 添加设备
type AddDeviceParam struct {
	DeviceName        string         `form:"deviceName" json:"deviceName" binding:"max=30"`
	DeviceIcon        string         `form:"deviceIcon" json:"deviceIcon" binding:"max=40"`
	DeviceGroup       string         `form:"deviceGroup" json:"deviceGroup" binding:"max=40"`
	DeviceSubjectName string         `form:"deviceSubjectName" json:"deviceSubjectName" binding:"max=40"`
	DeviceState       []model.Attrib `form:"deviceState" json:"deviceState" binding:"max=40"`
	DeviceControl     []model.Attrib `form:"deviceControl" json:"deviceControl" binding:"max=40"`
}

type GroupName struct {
	GroupName string `form:"groupName" json:"groupName" binding:"required"`
}

// 1,设备的注册 2,设备的订阅主题设置 3,设备的控制 4,设备的查询

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
func AddDevice(ctx *gin.Context) {
	//var addDeviceParam AddDeviceParam
	//if err := ctx.ShouldBind(&addDeviceParam); err != nil {
	//	result.APIResponse(ctx, code.ErrParam, err)
	//	return
	//}
	//// 表添加设备
	//err1 := db.AddDevice(&db.Device{
	//	DeviceID:    uuid.NewV4(),
	//	DeviceName:  addDeviceParam.DeviceName,
	//	DeviceGroup: addDeviceParam.DeviceGroup,
	//	DeviceSubject: db.DeviceSubject{
	//		DeviceState:   addDeviceParam.DeviceState,
	//		DeviceControl: addDeviceParam.DeviceControl,
	//	},
	//	DeviceTemplate: addDeviceParam.DeviceSubjectName,
	//})
	//if err1 != nil {
	//	result.APIResponse(ctx, code.ErrParam, err1)
	//	return
	//}
	//result.APIResponse(ctx, code.OK, "ok")
}
