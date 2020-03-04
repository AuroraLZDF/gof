/*
 * @Author: aurora
 * @Date: 2020-02-25 13:37:58
 * @LastEditors: aurora
 * @LastEditTime: 2020-02-28 08:23:39
 * @Description: 
 */
package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/auroraLZDF/gof/app/controller/welcome"
	"github.com/auroraLZDF/gof/app/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	/*r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)*/


	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello world",
		})
	})

	v1 := r.Group("/api/v1")
	v1.Use(middleware.JWT())
	{
		v1.GET("/", welcome.Index)

	}

	return r
}
