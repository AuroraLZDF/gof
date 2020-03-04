package test

import (
	"fmt"
	"github.com/auroraLZDF/gof/app/models"
	"github.com/auroraLZDF/gof/utils"
	"testing"
)

func TestUser( t *testing.T)  {
	//自动迁移
	err := DB.AutoMigrate(models.User{}).Error
	if err != nil {
		panic(err)
	}

	var u = models.User{UserName:"gaozhen",Password:utils.Md5("123456")}
	user, _ := u.UserCreate()
	fmt.Println("userInfo: ", user)
}
