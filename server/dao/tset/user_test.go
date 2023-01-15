package tset

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"linktree_server/modules/db/model"
	"linktree_server/server/dao"
	"testing"
	"time"
)

func TestUserCurd(t *testing.T) {
	sql := sqlite.Open("./test.db")
	db, err := gorm.Open(sql)
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	err1 := db.AutoMigrate(&model.User{})
	if err1 != nil {
		return
	}
	id := uuid.NewV4()
	userOld := model.User{
		UserID:     id,
		UserName:   "test",
		Tel:        123,
		Password:   "qdq",
		CreateTime: time.Now(),
		HeadUri:    "q",
		Role: model.RoleArray{
			"admin", "root",
		},
	}
	userNew := model.User{
		UserID:     id,
		UserName:   "test1",
		Tel:        123,
		Password:   "qdq",
		CreateTime: time.Now(),
		HeadUri:    "q",
		Role: model.RoleArray{
			"admin", "root",
		},
	}
	userDao := &dao.UserDao{
		Db: db,
	}
	// 添加
	t.Logf("添加=======>%+v", dao.User.DB(userDao).Create(userOld))
	// 添加
	t.Logf("查询所有=======>%+v", dao.User.DB(userDao).All())
	// 查询一个对象
	t.Logf("查询一个对象=======>%+v", dao.User.DB(userDao).Find(model.User{UserName: "test"}))
	// 更新
	t.Logf("更新=======>%+v", dao.User.DB(userDao).Update(id, userNew))
	// 删除
	t.Logf("删除=======>%+v", dao.User.DB(userDao).Delete(id))
}
