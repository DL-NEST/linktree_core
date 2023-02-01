package db

import (
	model2 "linktree_core/modules/database/db/model"
	"linktree_core/utils/glog"
)

//// 是否需要迁移
//func needMigration() bool {
//	return DB.Where("name = ?", "db_version_"+conf.RequiredDBVersion).First(&setting).Error != nil
//}

func migration() {

	glog.Log.Info("开始初始化数据库")
	// 自动迁移(表的创建)
	// 用户表
	err := DB.AutoMigrate(&model2.Device{})
	if err != nil {
		glog.Log.Error("用户表迁徙失败")
		return
	}

	err = DB.AutoMigrate(&model2.User{})
	if err != nil {
		glog.Log.Error("用户表迁徙失败")
		return
	}
}
