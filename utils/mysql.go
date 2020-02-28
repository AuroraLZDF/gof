package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/auroraLZDF/goconf"
	"time"
)

func DbInit(conf *goconf.Config) *gorm.DB {
	db := NewConn(conf)

	db.DB().SetMaxIdleConns(10)					//SetMaxIdleConns用于设置闲置的连接数
	db.DB().SetMaxOpenConns(100)					//SetMaxOpenConns用于设置最大打开的连接数
	db.DB().SetConnMaxLifetime(60 * time.Second)	// 防止出现 (invalid connection) 数据库连接超时错误


	db.LogMode(true)	// 启用Logger，显示详细日志

	// 自动迁移模式
	/*db.AutoMigrate(&Model.UserModel{},
		&Model.UserDetailModel{},
		&Model.UserAuthsModel{},
	)*/
	return db
}

func NewConn(conf *goconf.Config) *gorm.DB {
	format := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10ms",
		conf.GetValue("database", "username"),
		conf.GetValue("database", "password"),
		conf.GetValue("database", "hostname"),
		conf.GetValue("database", "port"),
		conf.GetValue("database", "dbname"),
		) + "&loc=Asia%2FShanghai"

	db, err := gorm.Open("mysql", format)
	if err != nil {
		panic("mysql connect error: " + err.Error())
	}

	return db
}
