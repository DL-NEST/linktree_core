package tests

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"linktree_core/server/dao"
	"linktree_core/server/model/entity"
	"testing"
)

func TestBaseUserDao(t *testing.T) {
	sql := sqlite.Open("./test.db")
	DB, err := gorm.Open(sql)
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	err1 := DB.AutoMigrate(&entity.User{})
	if err1 != nil {
		t.Errorf("Database migration failed: %v", err)
		return
	}
	userNew := entity.User{
		BaseModel: entity.BaseModel{},
		UUID:      uuid.New(),
		UserName:  "root",
		Tel:       12345,
		Password:  "admin",
		HeadUri:   "/head/{}.jpg",
	}
	// 设置db
	dao.User.DB(DB)

	// 添加
	err = dao.User.Create(&userNew)
	t.Logf("添加=======>%+v", err)
	// 查询表
	err, a := dao.User.All()
	t.Logf("查询所有=======>%+v", a)
	// 更新
	err, u := dao.User.Update(1, &entity.User{UserName: "admin"})
	t.Logf("更新=======>%+v", u)
	// 查询一个对象
	err, f := dao.User.Find(&entity.User{UserName: "admin"})
	t.Logf("查询一个对象=======>%+v", f)
	// 删除
	err = dao.User.DeleteUnscoped(1)
	t.Logf("删除=======>%+v", err)
}

func FuzzName(f *testing.F) {
	f.Fuzz(func(t *testing.T) {

	})
}
