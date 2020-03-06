package router

import (

	"github.com/gin-gonic/gin"
	//"io"
	"net/http"
	//"os"
	//"github.com/auroraLZDF/gof/app/middleware"
	"github.com/auroraLZDF/gof/app/controller/welcome"
	"github.com/auroraLZDF/gof/app/controller/user"

)

func InitRouter() *gin.Engine {
	// 写入日志的文件
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	/*
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	*/


	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello world",
		})
	})

	v1 := router.Group("/api/v1")
	v1.Use(/*middleware.Handler()*/)
	{
		v1.GET("/", welcome.Index)
		v1.GET("/user/list", user.Index)
		v1.GET("/user/create", user.Create)

	}

	return router
}
