package test

import (
	"fmt"
	"github.com/auroraLZDF/goconf"
	"github.com/auroraLZDF/gof/utils"
	"testing"
)

//var Conf *goconf.Config

func init() {
	Conf = goconf.InitConfig("./config/config.ini")
}

func TestConf(t *testing.T) {
	var currentPath = utils.GetCurrentDirectory()
	fmt.Println("current: " + currentPath)

	var path = "./config/config.ini"
	if !utils.IsFile(path) {
		fmt.Println("file not exists: " + path)
		return
	}
	Conf := goconf.InitConfig(path)
	Conf.DeleteValue("database", "username") //username 是你删除的 key

	username := Conf.GetValue("database", "username")
	if len(username) == 0 {
		fmt.Println("username is not exists") // this stdout username is not exists
	}
}
