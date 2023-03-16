package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"linktree_core/global"
)

// InitDBLink 创建数据库连接
func InitDBLink() {
	// 未初始化的时候跳过数据库连接
	if !global.State.SysInit {
		return
	}
	// 判断数据库的使用类型 sqlite and mysql
	if viper.GetString("persistence.db-type") == "mysql" {
		// 使用mysql
		global.GLOG.Info("\t连接数据库:mysql")
		global.DB = LinkMySql()
	} else {
		// 使用sqlite
		global.GLOG.Info("\t连接数据库:sqlite")
		global.DB = LinkSqlLite("./linkTree.sqlite")
	}
	// 数据库迁徙
	migration()
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
		viper.GetString("persistence.mysql.username"),
		viper.GetString("persistence.mysql.password"),
		viper.GetString("persistence.mysql.addr"),
		viper.GetString("persistence.mysql.dataname"),
		viper.GetString("persistence.mysql.charset"))
}
