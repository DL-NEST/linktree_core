package db

import (
	"linktree_core/modules/db/model"
	"linktree_core/utils/logger"
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
}
