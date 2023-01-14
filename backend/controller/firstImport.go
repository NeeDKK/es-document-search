package controller

import (
	"bufio"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/NeeDKK/esDocumentSearch/entity"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

type FirstImportController struct {
}

func FirstImport(c *gin.Context) {
	var path entity.ImportPath
	err := c.ShouldBindJSON(&path)
	if err != nil {
		fmt.Println("参数绑定失败:", err.Error())
	}
	//去除重复文件后再导入
	moveFile(path.Path)
	dir, err := ioutil.ReadDir(path.Path)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		file, err := os.Open(path.Path + fi.Name())
		if err != nil {
			fmt.Println(err)
		}
		//创建文件大小的字节数组
		stat, _ := file.Stat()
		bytes := make([]byte, stat.Size())
		buffer := bufio.NewReader(file)
		//将文件读取到字节数组中
		buffer.Read(bytes)
		//将字节数组转换为base64的字符串
		doc := base64.StdEncoding.EncodeToString(bytes)
		file.Close()
		FileToEs(c, doc)
	}
	entity.OkWithMessage("导入成功", c)
}

func moveFile(path string) {
	m := make(map[string]string)
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		file, err := os.Open(path + fi.Name())
		if err != nil {
			fmt.Println(err)
		}
		content, _ := ioutil.ReadAll(file)
		file.Close()
		md5h := md5.New()
		_, err = md5h.Write(content)
		if err != nil {
			fmt.Println(err)
		}
		//获取md5值在map中去重
		md5value := hex.EncodeToString(md5h.Sum(nil))
		if _, ok := m[md5value]; ok {
			fmt.Println("重复文件:", fi.Name())
			err := os.Rename(path+fi.Name(), path+"重复文件\\"+fi.Name())
			if err != nil {
				fmt.Println(err)
			}
		} else {
			m[md5value] = fi.Name()
		}
	}
}
