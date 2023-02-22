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
	FindInTel(tel uint) (error, entity.User)
	FindInUserName(userName string) (error, entity.User)
}

func (do userDao) FindInTel(tel uint) (error, entity.User) {
	var result entity.User
	err := do.DB().db.Where(&entity.User{Tel: tel}).First(&result).Error
	if err != nil {
		return err, result
	}
	return nil, result
}

func (do userDao) FindInUserName(userName string) (error, entity.User) {
	var result entity.User
	err := do.DB().db.Where(&entity.User{UserName: userName}).First(&result).Error
	if err != nil {
		return err, result
	}
	return nil, result
}
