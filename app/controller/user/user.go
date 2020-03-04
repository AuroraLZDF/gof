package user

import (
	"github.com/auroraLZDF/gof/app/models"
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
