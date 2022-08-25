/*
 * @Author: lxk20021217
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 16:43:13
 * @Description: 加载配置
 */

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

	if err := viper.ReadInConfig(); err != nil {
		logger.Error("读取配置失败: %v", err)
		return
	}

	if err := viper.Unmarshal(&CONFIG); err != nil {
		logger.Error("解析配置失败: %v", err)
		return
	}

	logger.Info("读取配置成功！")
}
