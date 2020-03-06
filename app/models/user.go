package models

import (
	"fmt"
	"github.com/auroraLZDF/gof/config"
	"log"
)

type User struct {
	Model
	UserName	string 	`gorm:"not null;default:'';comment:'用户名'"`
	Password	string 	`gorm:"not null;default:'';comment:'密码'"`
	Avatar		string 	`gorm:"not null;default:'';comment:'头像'"`
	Email 		string 	`gorm:"not null;default:'';comment:'手机号'"`
	Mobile 		string 	`gorm:"not null;default:'';comment:'手机号'"`
	Sex 		int 	`gorm:"not null;default:0;comment:'性别：1=》男；2=》女；0=》未知'"`
	Status      int     `gorm:"not null;default:0;comment:'登录状态：0=》初始；1=》允许登录；2=》禁止登录'"`
}

func (User) TableName() string {
	return "user"
}

func (u User) UserInfo(field string, value string) (user User, err error) {
	err = config.Db.Where(field + "= ?", value).First(&user).Error
	if err != nil {
		log.Println("userInfo: " + err.Error())
	}
	return
}

func (u User) UserList(params map[string]interface{}) (users []User, err error) {
	log.Println("db:", config.Db, "params", params)
	err = config.Db.Where(params).Find(&users).Error	//db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	if err != nil {
		log.Println("userList: " + err.Error())
	}
	return
}

func (u User) UserCreate() (user User, err error) {
	fmt.Println("user:", u)
	err = config.Db.Create(&u).Error
	if err != nil {
		log.Println("userCreate: " + err.Error())
	}

	user = u
	return
}