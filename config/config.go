package config

import (
	"encoding/json"
	"github.com/auroraLZDF/gof/utils"
	"io/ioutil"
	"log"
)

/*const (
	TimeFormat = "2006-01-02 15:04:05"
)*/

type config struct {
	AppPath string `json:"app_path"`

	DbHost string `json:"db_host"`
	DbUser string `json:"db_user"`
	DbPass string `json:"db_pass"`
	DbPort string `json:"db_port"`
	DbName string `json:"db_name"`

	Env             string `json:"env";default:"production"`
}

var Config config

func SetConfig() {
	path := utils.CurrentPath() + "/../.env.json"
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic("读取配置文件失败")
	}

	err = json.Unmarshal(file, &Config)
	if err != nil {
		log.Fatal("[loadConfig]: %s\n", err.Error())
	}
}
