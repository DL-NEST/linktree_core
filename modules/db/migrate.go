package db

import (
	"linktree_server/modules/db/model"
	"linktree_server/utils/logger"
)

//// 是否需要迁移
//func needMigration() bool {
//	return DB.Where("name = ?", "db_version_"+conf.RequiredDBVersion).First(&setting).Error != nil
//}

func migration() {

	logger.Log.Info("开始初始化数据库")
	// 自动迁移(表的创建)
	// 用户表
	err := DB.AutoMigrate(&model.Device{})
	if err != nil {
		logger.Log.Error("用户表迁徙失败")
		return
	}

	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		logger.Log.Error("用户表迁徙失败")
		return
	}

	err = DB.AutoMigrate(&model.Test{})
	if err != nil {
		logger.Log.Error("用户表迁徙失败")
		return
	}

	//err1 := AddUser(&User{
	//	UserID:     uuid.NewV4(),
	//	UserName:   "root",
	//	Tel:        234252,
	//	Password:   "qq2002123",
	//	CreateTime: time.Now(),
	//	HeadUri:    "",
	//	Role: RoleArray{
	//		"ssaf", "asef",
	//	},
	//})
	//if err1 != nil {
	//	return
	//}

	//logger.Log.Info(FindDeviceInName("台灯").DeviceSubject.DeviceState[0].AttribName)

}
