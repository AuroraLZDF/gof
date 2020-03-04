package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/auroraLZDF/gof/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = utils.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = utils.INVALID_PARAMS
		}

		if code != utils.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  utils.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
