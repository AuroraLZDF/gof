package test

import (
	"fmt"
	"github.com/auroraLZDF/goconf"
	"github.com/auroraLZDF/gof/app/models"
	"github.com/auroraLZDF/gof/utils"
	"github.com/jinzhu/gorm"
	"testing"
)

var DB *gorm.DB

func TestUser( t *testing.T)  {
	conf := goconf.InitConfig("../config/config.ini")
	DB = utils.DbInit(conf)

	//自动迁移
	err := DB.AutoMigrate(models.User{}).Error
	if err != nil {
		panic(err)
	}

	fmt.Println("autoMigrate success!")

	//var u = models.User{UserName:"gaozhen",Password:utils.Md5("123456")}
	//user, _ := u.UserCreate()
	//fmt.Println("userInfo: ", user)
}
