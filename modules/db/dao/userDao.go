package dao

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"linktree_server/modules/db/model"
)

// IUserDao 定义一个接口
type IUserDao interface {
	DB(db *gorm.DB) IUserDao
	All() []model.User
	Find(v model.User) []model.User
	Create(v model.User) model.User
	Update(pKey uuid.UUID, v model.User) model.User
	Delete(Key uuid.UUID) error
}

// UserDao 实现了IDao接口
type UserDao struct {
	db *gorm.DB
}

func (do UserDao) DB(db *gorm.DB) IUserDao {
	do.db = db
	return do
}
func (do UserDao) All() []model.User {
	var list []model.User
	do.db.Find(&list)
	return list
}
func (do UserDao) Find(v model.User) []model.User {
	var list []model.User
	do.db.Where(&v).Find(&list)
	return list
}
func (do UserDao) Create(v model.User) model.User {
	do.db.Create(&v)
	return v
}
func (do UserDao) Update(pKey uuid.UUID, v model.User) model.User {
	var user model.User
	do.db.First(&user, pKey).Updates(&v)
	return v
}
func (do UserDao) Delete(pKey uuid.UUID) error {
	return do.db.Delete(model.User{UserID: pKey}).Error
}
