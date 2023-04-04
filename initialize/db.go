package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"linktree_core/global"
	"linktree_core/server/model/entity"
)

const (
	Mysql  = "mysql"
	Sqlite = "sqlite"
)

// InitDBLink 创建数据库连接
func InitDBLink() {
	// 未初始化的时候跳过数据库连接
	if !global.State.SysInit {
		return
	}
	switch global.Config.GetString("persistence.db-type") {
	case Mysql:
		LinkMySql()
	case Sqlite:
		LinkSqlLite("./linkTree.sqlite")
	default:
	}
	// 判断数据库的使用类型 sqlite and mysql
	if global.Config.GetString("persistence.db-type") == "mysql" {
		// 使用mysql
		global.GLOG.Info("连接数据库:mysql")
		global.DB = LinkMySql()
	} else {
		// 使用sqlite
		global.GLOG.Info("连接数据库:sqlite")
		global.DB = LinkSqlLite("./linkTree.sqlite")
	}
	// 数据库迁徙
	RegisterTables()
}

func LinkMySql() *gorm.DB {
	// 打开数据库
	db, err := gorm.Open(mysql.Open(Dsn()), &gorm.Config{
		CreateBatchSize: 1000,
	})
	if err != nil {
		global.GLOG.Panic("连接数据库失败")
	}
	return db
}

func LinkSqlLite(path string) *gorm.DB {
	// 打开数据库
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		CreateBatchSize: 1000,
	})
	if err != nil {
		global.GLOG.Panic("连接数据库失败")
	}
	return db
}

func Dsn() string {
	// "username:password@tcp(ip:port)database_name?charset=utf8"
	return fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",
		global.Config.GetString("persistence.mysql.username"),
		global.Config.GetString("persistence.mysql.password"),
		global.Config.GetString("persistence.mysql.addr"),
		global.Config.GetString("persistence.mysql.dataname"),
		global.Config.GetString("persistence.mysql.charset"))
}

func RegisterTables() {
	// 自动迁移
	err := global.DB.AutoMigrate(
		&entity.Device{},
		&entity.User{},
		&entity.DeviceMsg{},
	)
	if err != nil {
		global.GLOG.Error("表迁徙失败")
		return
	}
}
