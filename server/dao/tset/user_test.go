package tset

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"linktree_core/modules/database/db/model"
	"linktree_core/server/dao"
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
	id := uuid.New()
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
	err, c := dao.User.DB(userDao).Create(userOld)
	t.Logf("添加=======>%+v", c)
	// 添加
	err, a := dao.User.DB(userDao).All()
	t.Logf("查询所有=======>%+v", a)
	// 查询一个对象
	err, f := dao.User.DB(userDao).Find(model.User{UserName: "test"})
	t.Logf("查询一个对象=======>%+v", f)
	// 更新
	err, u := dao.User.DB(userDao).Update(id, userNew)
	t.Logf("更新=======>%+v", u)
	// 删除
	err = dao.User.DB(userDao).Delete(id)
	t.Logf("删除=======>%+v", err)
}
