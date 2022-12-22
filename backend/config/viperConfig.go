package config

import (
	"fmt"
	"github.com/NeeDKK/esDocumentSearch/entity"
	"github.com/spf13/viper"
)

var GlobalConfig *entity.Config

func InitViper() {
	// 设置配置文件的名字
	viper.SetConfigName("config")
	// 设置配置文件所在的路径
	viper.AddConfigPath(".")
	// 设置配置文件类型
	viper.SetConfigType("yaml")
	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件失败:", err.Error())
		panic(err)
	}
	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		fmt.Println("读取配置文件失败:", err.Error())
		panic(err)
	}
}
