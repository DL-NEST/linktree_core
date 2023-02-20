package db

import (
	"linktree_core/global"
	entity2 "linktree_core/server/model/entity"
)

func migration() {
	// 自动迁移(表的创建)
	// 用户表
	err := global.DB.AutoMigrate(&entity2.Device{})
	if err != nil {
		global.GLOG.Error("用户表迁徙失败")
		return
	}

	err = global.DB.AutoMigrate(&entity2.User{})
	if err != nil {
		global.GLOG.Error("用户表迁徙失败")
		return
	}
}
