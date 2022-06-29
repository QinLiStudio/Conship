//Author: lxk20021217
//Date: 2022-06-19 00:33:25
//LastEditTime: 2022-06-19 19:38:53
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\db\redis.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-19 00:33:25
//LastEditTime: 2022-06-19 19:41:55
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\db\redis.go
//是谁总是天亮了才睡

package db

import (
	"log"

	"github.com/QinLiStudio/Conship/internal/app/configs"
	"github.com/go-redis/redis"
)

func Redis() {
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
		log.Fatal("连接redis失败: ", err)
	} else {
		configs.REDISDB = client
		log.Print("连接redis成功: ", pong)
	}
}
