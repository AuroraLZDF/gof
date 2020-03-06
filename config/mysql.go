package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"time"
)

var Db *gorm.DB

//初始化方法
func SetDb() {
	Db = NewConn()
	Db.DB().SetMaxIdleConns(10)  // 数据库的空闲连接
	Db.DB().SetMaxOpenConns(100) // 数据库的最大连接
	Db.DB().SetConnMaxLifetime(60 * time.Second)	// 防止出现 (invalid connection) 数据库连接超时错误

	if env := Config.Env; env != "production" {
		Db.LogMode(true)	// 启用Logger，显示详细日志
	}
}

func NewConn() *gorm.DB {
	format := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10ms",
		Config.DbUser, Config.DbPass, Config.DbHost, Config.DbPort, Config.DbName) + "&loc=Asia%2FShanghai"

	db, err := gorm.Open("mysql", format)
	if err != nil {
		panic("mysql connect error: " + err.Error())
	}

	return db
}
