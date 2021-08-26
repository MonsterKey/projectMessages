package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"黄泽林是": "sbKing",
		})
	})
	r.GET("/test01/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK,"Hello %s",name)
	})
	r.GET("/test01/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	r.GET("/test02", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname","黄")
		lastname := context.Query("lastname")
		context.String(http.StatusOK, "Hello, %s%s",firstname, lastname)
	})
	r.POST("/testpost01", func(context *gin.Context) {
		message := context.DefaultPostForm("message","黄泽林是傻逼！")
		nick := context.PostForm("nick")
		getmessage := context.Query("getmessage")
		getnick := context.DefaultQuery("getnick","getnick")

		context.JSON(200, gin.H{
			"message": message,
			"nick": nick,
			"status":  "posted",
			"getmessage": getmessage,
			"getnick": getnick,
		})
	})
	r.POST("/upload", func(context *gin.Context) {
		file,error := context.FormFile("file")
		log.Println(file)
		if error != nil{
			context.String(http.StatusBadRequest, "请求失败")
			return
		}
		//上传到指定路径
		//dst := "/Users/ironmr/Downloads/environment/image"
		//context.SaveUploadedFile(file,dst)
		filename := file.Filename
		dst := "/Users/ironmr/Downloads/environment/image/"+filename
		if err := context.SaveUploadedFile(file, dst); err != nil {
			context.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
			return
		}

		context.String(http.StatusOK, fmt.Sprintf("'%s' upload=====图片上传成功！",file.Filename))
	})
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}