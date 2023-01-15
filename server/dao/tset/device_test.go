package tset

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"linktree_server/modules/db/model"
	"linktree_server/server/dao"
	"testing"
)

func TestDeviceCurd(t *testing.T) {
	sql := sqlite.Open("./test.db")
	db, err := gorm.Open(sql)
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	err1 := db.AutoMigrate(&model.Device{})
	if err1 != nil {
		return
	}
	id := uuid.NewV4()
	DeviceOld := model.Device{
		DeviceID:    id,
		DeviceName:  "智能设备",
		DeviceGroup: "公共",
		DeviceSubject: model.DeviceSubject{
			DeviceState: []model.Attrib{
				{
					AttribName:  "温度",
					AttribParse: "int",
				},
			},
			DeviceControl: []model.Attrib{
				{
					AttribName:  "开关",
					AttribParse: "bool",
				},
			},
		},
		DeviceTemplate: "默认模板",
	}
	DeviceNew := model.Device{
		DeviceID:    id,
		DeviceName:  "智能设备",
		DeviceGroup: "公共",
		DeviceSubject: model.DeviceSubject{
			DeviceState: []model.Attrib{
				{
					AttribName:  "温度",
					AttribParse: "int",
				},
			},
			DeviceControl: []model.Attrib{
				{
					AttribName:  "开关",
					AttribParse: "bool",
				},
			},
		},
		DeviceTemplate: "默认模板",
	}
	DeviceDao := &dao.DeviceDao{
		Db: db,
	}
	// 添加
	t.Logf("添加=======>%+v", dao.Device.DB(DeviceDao).Create(DeviceOld))
	// 添加
	t.Logf("查询所有=======>%+v", dao.Device.DB(DeviceDao).All())
	// 查询一个对象
	t.Logf("查询一个对象=======>%+v", dao.Device.DB(DeviceDao).Find(model.Device{DeviceName: "智能设备"}))
	// 更新
	t.Logf("更新=======>%+v", dao.Device.DB(DeviceDao).Update(id, DeviceNew))
	// 删除
	t.Logf("删除=======>%+v", dao.Device.DB(DeviceDao).Delete(id))
}
