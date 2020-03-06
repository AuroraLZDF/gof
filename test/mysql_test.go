package test

import (
	"fmt"
	"github.com/auroraLZDF/goconf"
	"github.com/auroraLZDF/gof/utils"
	"testing"
)

func TestMysqlConn( t *testing.T)  {
	var conf *goconf.Config
	conf = goconf.InitConfig("../config/config.ini")

	db := utils.NewConn(conf)
	if db.Error != nil {
		fmt.Println(db.Error)
	}
	fmt.Println("mysql connect success!")
}



