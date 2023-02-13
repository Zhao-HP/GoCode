package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var GlobalConfig *viper.Viper

func init() {
	fmt.Println("读取全局配置")
	GlobalConfig = viper.New()
	GlobalConfig.AddConfigPath("/Users/zhao/Desktop/Code/Go/GoCode/web/gin/config")
	GlobalConfig.SetConfigName("config")
	GlobalConfig.SetConfigType("yml")
	if err := GlobalConfig.ReadInConfig(); err != nil {
		fmt.Println("读取全局配置异常: ", err)
		return
	}

	initDbConfig()
	initLogger()
	initRedis()
}
