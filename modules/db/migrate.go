package db

import (
	"linktree_core/server/entity"
	"linktree_core/utils/dlog"
)

//// 是否需要迁移
//func needMigration() bool {
//	return DB.Where("name = ?", "db_version_"+conf.RequiredDBVersion).First(&setting).Error != nil
//}

func migration() {

	dlog.Log.Info("开始初始化数据库")
	// 自动迁移(表的创建)
	// 用户表
	err := DB.AutoMigrate(&entity.Device{})
	if err != nil {
		dlog.Log.Error("用户表迁徙失败")
		return
	}

	err = DB.AutoMigrate(&entity.User{})
	if err != nil {
		dlog.Log.Error("用户表迁徙失败")
		return
	}
}