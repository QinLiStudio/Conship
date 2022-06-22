//Author: lxk20021217
//Date: 2022-06-19 00:33:25
//LastEditTime: 2022-06-19 00:41:27
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\DB\redis.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-19 00:33:25
//LastEditTime: 2022-06-19 00:46:05
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\DB\redis.go
//是谁总是天亮了才睡

package db

import (
	"log"

	"github.com/QinLiStudio/Conhip/internal/app/configs"
	"github.com/go-redis/redis"
)

func initRedis() {
	// 获取 rdis 配置
	r := configs.CONFIG.Redis
	// 建立连接
	client := redis.NewClint(&redis.Options{
		Addr:     r.Add,
		Pssword: r.Password,
		DB:      r.DB,
	})
	// 测试连接
	pong, err := client.Ping().Result()
	if err ! nil {
		log.Fatal("Connect redisfailed: ", err)
	} else {
		lobal.REDIS_DB = client
	log.Print("Connect redis success: ", pong)
	}
}
