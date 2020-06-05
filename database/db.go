package database

import (
	"log"

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
	Eloquent, DbErr = gorm.Open("sqlite3", "/home/sun/Works/gopath/src/newapp/movie.sqlite")
	if DbErr != nil {
		log.Fatal("error daabase")
	}
	// 全局禁用表名复数
	Eloquent.SingularTable(true)
	// 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
}
