/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-06-19 16:59:39
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-07-08 18:16:35
 * @FilePath: \Conship\internal\app\redis.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */
package app

import (
	"log"

	refraction1 "github.com/QinLiStudio/Conship/pkg/utils/refraction_struct"
	"github.com/go-redis/redis"
)

func Redis() {

	// 获取 redis 配置
	r := refraction1.ToRedis
	// 建立连接
	client := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
	})
	// 测试连接
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal("Connect redis failed: ", err)
	} else {
		DBR = client
		log.Print("Connect redis success: ", pong)
	}
}
