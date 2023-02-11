package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"linktree_core/utils/dlog"
)

var DB *gorm.DB

// CreateDBLink 创建数据库连接
func CreateDBLink() {

	// 判断数据库的使用类型 sqlite and mysql
	if viper.InConfig("persistence.mysql.addr") {
		// 使用mysql
		dlog.Log.Info("\t连接数据库:mysql")
		DB = LinkMySql()
	} else {
		// 使用sqlite
		dlog.Log.Info("\t连接数据库:sqlite")
		DB = LinkSqlLite("./linkTree.sqlite")
	}
	// 数据库迁徙
	migration()
}

func LinkMySql() *gorm.DB {
	// "username:password@tcp(ip:port)database_name?charset=utf8"
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("persistence.mysql.username"),
		viper.GetString("persistence.mysql.password"),
		viper.GetString("persistence.mysql.addr"),
		viper.GetString("persistence.mysql.dataname"),
		viper.GetString("persistence.mysql.charset"))
	// 打开数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 1000,
	})
	if err != nil {
		panic(any("连接数据库失败"))
	}
	DB = db
	return db
}

func LinkSqlLite(path string) *gorm.DB {
	// 打开数据库
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		CreateBatchSize: 1000,
	})
	if err != nil {
		panic(any("连接数据库失败"))
	}
	DB = db
	return db
}
