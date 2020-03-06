package user

import (
	"github.com/auroraLZDF/gof/app/models"
	"github.com/auroraLZDF/gof/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	var U = models.User{}
	lists, _ := U.UserList(map[string]interface{}{"status": 1})


	c.JSON(http.StatusOK, gin.H{
		"data": lists,
		"msg" : "ok",
	})
}

func Create(c *gin.Context)  {
	var user = models.User{
		UserName: c.PostForm("username"),
		Password: utils.Md5(c.PostForm("password")),
		Avatar: c.DefaultPostForm("avatar", ""),
		Email: c.PostForm("email"),
		Mobile: c.PostForm("mobile"),
		Sex: utils.StringToInt(c.PostForm("sex")),
		Status: 0,
	}

	_user, err := user.UserCreate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": "",
			"msg" : "创建用户失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": _user,
		"msg": "创建用户成功",
	})
}
