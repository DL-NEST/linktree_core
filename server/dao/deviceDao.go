package dao

import (
	"linktree_core/server/entity"
)

// deviceDao object
type deviceDao struct {
	baseDao[entity.User]
}

// IDeviceDao interface
type IDeviceDao interface {
	FindInUserName()
}

func (do userDao) FindInDeviceName() {
	//TODO implement me
	panic("implement me")
}
