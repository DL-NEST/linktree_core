package dao

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"linktree_server/modules/db"
	"linktree_server/modules/db/model"
)

// IUserDao interface
type IUserDao interface {
	DB(UserOpt ...UserOpt) IUserDao
	All() []model.User
	Find(v model.User) []model.User
	Create(v model.User) model.User
	Update(pKey uuid.UUID, v model.User) model.User
	Delete(Key uuid.UUID) error
}

// UserDao object
type UserDao struct {
	Db *gorm.DB
}

// Apply to object
func (do *UserDao) Apply(config *UserDao) {
	*config = *do
}

// UserOpt interface
type UserOpt interface {
	Apply(*UserDao)
}

// DB 设置db句柄来源
func (do *UserDao) DB(UserOpt ...UserOpt) IUserDao {
	var userDao UserDao
	for _, o := range UserOpt {
		o.Apply(&userDao)
	}
	if userDao.Db != nil {
		do.Db = userDao.Db
		return do
	}
	do.Db = db.DB
	return do
}
func (do *UserDao) All() []model.User {
	var list []model.User
	do.Db.Find(&list)
	return list
}
func (do *UserDao) Find(v model.User) []model.User {
	var list []model.User
	do.Db.Where(&v).Find(&list)
	return list
}
func (do *UserDao) Create(v model.User) model.User {
	do.Db.Create(&v)
	return v
}
func (do *UserDao) Update(pKey uuid.UUID, v model.User) model.User {
	var user model.User
	do.Db.First(&user, pKey).Updates(&v)
	return v
}
func (do *UserDao) Delete(pKey uuid.UUID) error {
	return do.Db.Delete(model.User{UserID: pKey}).Error
}
