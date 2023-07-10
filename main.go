package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	// viper: 用来处理程序配置的
	// 环境变量加载我们的程序配置
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("DB_USERNAME", "dukuan")
	viper.SetDefault("DB_PASSWORD", "dotbalo")
	viper.SetDefault("DB_ADDRESS", "127.0.0.1")
	viper.SetDefault("DB_PORT", 3306)
	// 获取环境变量的配置
	// export DB_USERNAME=dddddddd
	viper.AutomaticEnv()
	logLevel := viper.GetString("LOG_LEVEL") // 获取程序的配置 --> debug
	fmt.Println(logLevel)

}
