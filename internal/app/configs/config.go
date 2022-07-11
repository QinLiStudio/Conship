//Author: lxk20021217
//Date: 2022-06-17 16:28:47
//LastEditTime: 2022-06-30 00:55:58
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\configs\config.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-17 16:28:47
//LastEditTime: 2022-07-10 15:39:03
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\configs\config.go
//是谁总是天亮了才睡

package configs

import (
	"github.com/QinLiStudio/Conship/pkg/logger"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	POSTGRESDB *gorm.DB
	REDISDB    *redis.Client
	CONFIG     Config
)

func InitConfig() {

	viper.SetConfigFile("./configs/config.toml")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Error("读取配置失败: %v", err)
		return
	}

	if err := viper.Unmarshal(&CONFIG); err != nil {
		logger.Error("解析配置失败: %v", err)
		return
	}

	logger.Info("读取配置成功！")
	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	// 配置文件发生变更之后会调用的回调函数
	//	fmt.Println("配置文件更改:", e.Name)
	//})
}
