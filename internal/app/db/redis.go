/*
 * @Author: lxk20021217
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 16:57:35
 * @Description: Redis 连接
 */

package db

import (
	"github.com/QinLiStudio/Conship/internal/app/configs"
	"github.com/QinLiStudio/Conship/pkg/logger"
	"github.com/go-redis/redis"
)

func InitRedis() {

	// 获取 redis 配置
	r := configs.CONFIG.Redis

	// 建立连接
	client := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
	})

	// 测试连接
	pong, err := client.Ping().Result()

	if err != nil {

		logger.Error("连接 redis 失败: %v", err)

	} else {

		configs.REDISDB = client

		logger.Info("连接 redis 成功: %v", pong)

	}
}
