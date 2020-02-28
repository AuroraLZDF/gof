package main

import (
	. "github.com/auroraLZDF/gof/router"
	"github.com/auroraLZDF/gof/utils"
	"github.com/jinzhu/gorm"
	"github.com/auroraLZDF/goconf"
)

var DB *gorm.DB
var Conf *goconf.Config

func init() {
	Conf = goconf.InitConfig("./config/config.ini")
	DB = utils.DbInit(Conf)

}

func main() {

	//当整个程序完成之后关闭数据库连接
	defer DB.Close()

	router := InitRouter()
	router.Run(":8080")
}
