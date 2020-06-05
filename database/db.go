package database

import (
	"log"

	utils "newapp/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //加载mysql
)

// Eloquent is a db connent
var (
	Eloquent *gorm.DB
	DbErr    error
)

func init() {
	// Eloquent, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local&timeout=10ms")
	// utils.GetDBPath 获取程序运行目录 返回sqlite数据库文件.第二个参数是dev模式true表示dev返回预定目录
	Eloquent, DbErr = gorm.Open("sqlite3", utils.GetDBPath("d", true))
	if DbErr != nil {
		log.Fatal("error daabase")
	}
	Eloquent.LogMode(true)
	// 全局禁用表名复数
	Eloquent.SingularTable(true)
	// 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
}
