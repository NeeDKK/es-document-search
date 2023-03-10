package main

import (
	"github.com/NeeDKK/esDocumentSearch/config"
	"github.com/NeeDKK/esDocumentSearch/controller"
	"github.com/gin-gonic/gin"
	"strconv"
)

func init() {
	// 初始化配置文件
	config.InitViper()
	// 初始化es连接
	config.InitEs()
	//读取文件构建本地缓存
	config.ReadFileForMap()
}

func main() {
	//创建路由
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	//配置跨域
	engine.Use(config.Cors())
	engine.GET("/search", controller.Search)
	engine.POST("/uploadFile", controller.UploadFile)
	engine.POST("/firstImport", controller.FirstImport)
	//启动gin服务
	engine.Run(":" + strconv.Itoa(config.GlobalConfig.Server.Port))
}
