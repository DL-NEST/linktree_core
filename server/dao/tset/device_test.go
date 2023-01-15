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
	err, c := dao.Device.DB(DeviceDao).Create(DeviceOld)
	t.Logf("添加=======>%+v", c)
	// 添加
	err, a := dao.Device.DB(DeviceDao).All()
	t.Logf("查询所有=======>%+v", a)
	// 查询一个对象
	err, f := dao.Device.DB(DeviceDao).Find(model.Device{DeviceName: "智能设备"})
	t.Logf("查询一个对象=======>%+v", f)
	// 更新
	err, u := dao.Device.DB(DeviceDao).Update(id, DeviceNew)
	t.Logf("更新=======>%+v", u)
	// 删除
	err = dao.Device.DB(DeviceDao).Delete(id)
	t.Logf("删除=======>%+v", err)
}
