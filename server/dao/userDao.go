package dao

import (
	"linktree_core/server/model/entity"
)

// userDao object
type userDao struct {
	baseDao[entity.User]
}

// IUserDao interface
type IUserDao interface {
	FindInUserName()
}

func (do userDao) FindInUserName() {

	//TODO implement me
	panic("implement me")
}
