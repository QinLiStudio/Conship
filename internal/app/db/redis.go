//Author: lxk20021217
//Date: 2022-06-19 00:33:25
//LastEditTime: 2022-06-19 19:38:53
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\db\redis.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-19 00:33:25
//LastEditTime: 2022-07-10 15:35:54
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\db\redis.go
//是谁总是天亮了才睡

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
