//Author: lxk20021217
//Date: 2022-06-17 16:28:47
//LastEditTime: 2022-06-19 00:34:52
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\configs\config.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-17 16:28:47
//LastEditTime: 2022-06-19 00:42:05
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\configs\config.go
//是谁总是天亮了才睡

package configs

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	MYSQLDB *gorm.DB
	REDISDB *redis.Client
	CONFIG  Config
)

func initConfig() {

	viper.SetConfigFile("./configs/config.toml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("读取配置失败: ", err)
	}

	if err := viper.Unmarshal(&CONFIG); err != nil {
		log.Fatal("解析配置失败: ", err)
	}

	viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	// 配置文件发生变更之后会调用的回调函数
	//	fmt.Println("配置文件更改:", e.Name)
	//})
}
