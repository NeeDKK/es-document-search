package controller

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/NeeDKK/esDocumentSearch/config"
	"github.com/NeeDKK/esDocumentSearch/entity"
	"github.com/gin-gonic/gin"
	"math/rand"
	"path"
	"strconv"
	"strings"
	"time"
)

type UploadController struct {
}

func Upload(c *gin.Context) {
	var resume entity.Resume
	err := c.ShouldBindJSON(&resume)
	if err != nil {
		fmt.Println("参数绑定失败:", err.Error())
	}
	sEnc := base64.StdEncoding.EncodeToString([]byte(resume.Content))
	resume.Content = sEnc
	do, err := config.EsClient.Index().Index(config.RESUMEINDEX).BodyJson(resume).Pipeline(config.RESUMEPIPLINE).Do(c)
	if err != nil {
		fmt.Println(err.Error())
		entity.FailWithDetailed(err, "上传失败", c)
		return
	}
	entity.OkWithDetailed(do, "上传成功", c)
}

func UploadFile(c *gin.Context) {
	_, file, err := c.Request.FormFile("file")
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名
	name := strings.TrimSuffix(file.Filename, ext)
	f, openError := file.Open() // 读取文件
	if openError != nil {
		fmt.Println("function file.Open() Filed", openError.Error())
		entity.FailWithDetailed(openError, "上传文件"+name+"失败", c)
		return
	}
	defer f.Close() // 创建文件 defer 关闭

	if err != nil {
		fmt.Println("文件上传失败:", err.Error())
	}
	//创建文件大小的字节数组
	bytes := make([]byte, file.Size)
	buffer := bufio.NewReader(f)
	//将文件读取到字节数组中
	buffer.Read(bytes)
	//将字节数组转换为base64的字符串
	doc := base64.StdEncoding.EncodeToString(bytes)
	var resume entity.Resume
	rand.Seed(time.Now().UnixNano())
	//正常从数据库读取
	resume.ID = uint(rand.Intn(9999))
	//录入简历数据时，从用户输入读取
	resume.Name = "name:" + strconv.Itoa(int(resume.ID))
	resume.Content = doc
	//定义索引和pipline写入es
	do, err := config.EsClient.Index().Index(config.RESUMEINDEX).BodyJson(resume).Pipeline(config.RESUMEPIPLINE).Do(c)
	if err != nil {
		fmt.Println(err.Error())
		entity.FailWithDetailed(err, "上传失败", c)
		return
	}
	entity.OkWithDetailed(do, "上传成功", c)
}
