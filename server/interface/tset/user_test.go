package tset

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_interface "linktree_core/server/interface"
	entity2 "linktree_core/server/model/entity"
	"testing"
)

func TestBaseUserDao(t *testing.T) {
	sql := sqlite.Open("./test.db")
	DB, err := gorm.Open(sql)
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	err1 := DB.AutoMigrate(&entity2.User{})
	if err1 != nil {
		t.Errorf("Database migration failed: %v", err)
		return
	}
	userNew := entity2.User{
		BaseModel: entity2.BaseModel{},
		UserID:    uuid.New(),
		UserName:  "root",
		Tel:       12345,
		Password:  "admin",
		HeadUri:   "/head/{}.jpg",
	}
	// 设置db
	_interface.User.DB(DB)

	// 添加
	err = _interface.User.Create(&userNew)
	t.Logf("添加=======>%+v", err)
	// 查询表
	err, a := _interface.User.All()
	t.Logf("查询所有=======>%+v", a)
	// 更新
	err, u := _interface.User.Update(1, &entity2.User{UserName: "admin"})
	t.Logf("更新=======>%+v", u)
	// 查询一个对象
	err, f := _interface.User.Find(&entity2.User{UserName: "admin"})
	t.Logf("查询一个对象=======>%+v", f)
	// 删除
	err = _interface.User.DeleteUnscoped(1)
	t.Logf("删除=======>%+v", err)
}

func FuzzName(f *testing.F) {
	f.Fuzz(func(t *testing.T) {

	})
}
