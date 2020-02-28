/*
 * @Author: aurora
 * @Date: 2020-02-25 09:09:17
 * @LastEditors: aurora
 * @LastEditTime: 2020-02-28 08:23:23
 * @Description:
 */
package welcome

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "api v1: hello world",
	})
}
