package dao

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"linktree_server/modules/db"
	"linktree_server/modules/db/model"
)

// IDeviceDao interface
type IDeviceDao interface {
	DB(option ...DeviceOpt) IDeviceDao
	All() []model.Device
	Find(v model.Device) []model.Device
	Create(v model.Device) model.Device
	Update(pKey uuid.UUID, v model.Device) model.Device
	Delete(Key uuid.UUID) error
}

// DeviceDao object
type DeviceDao struct {
	Db *gorm.DB
}

// Apply to object
func (do *DeviceDao) Apply(config *DeviceDao) {
	*config = *do
}

// DeviceOpt interface
type DeviceOpt interface {
	Apply(*DeviceDao)
}

// DB 设置db句柄来源
func (do *DeviceDao) DB(option ...DeviceOpt) IDeviceDao {
	var DeviceDao DeviceDao
	for _, o := range option {
		o.Apply(&DeviceDao)
	}
	if DeviceDao.Db != nil {
		do.Db = DeviceDao.Db
		return do
	}
	do.Db = db.DB
	return do
}
func (do *DeviceDao) All() []model.Device {
	var list []model.Device
	do.Db.Find(&list)
	return list
}
func (do *DeviceDao) Find(v model.Device) []model.Device {
	var list []model.Device
	do.Db.Where(&v).Find(&list)
	return list
}
func (do *DeviceDao) Create(v model.Device) model.Device {
	do.Db.Create(&v)
	return v
}
func (do *DeviceDao) Update(pKey uuid.UUID, v model.Device) model.Device {
	var Device model.Device
	do.Db.First(&Device, pKey).Updates(&v)
	return v
}
func (do *DeviceDao) Delete(pKey uuid.UUID) error {
	return do.Db.Delete(model.Device{DeviceID: pKey}).Error
}
