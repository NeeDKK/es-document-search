package main

import (
	"esDocumentSearch/config"
	"esDocumentSearch/controller"
	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化es连接
	config.InitEs()
}

func main() {
	//创建路由
	engine := gin.Default()
	engine.Use(config.Cors())
	engine.POST("/upload", controller.Upload)
	engine.GET("/search", controller.Search)
	engine.POST("/uploadFile", controller.UploadFile)
	//启动gin服务
	engine.Run(":9999")
}
