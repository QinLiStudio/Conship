/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-06-19 16:59:39
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-07-08 22:15:51
 * @FilePath: \Conship\internal\app\app.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */
package app

import (
	"database/sql"

	"github.com/QinLiStudio/Conship/pkg/utils/refraction_struct"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

// type Global struct {
// 	DBC      *refraction1.DBconf //总数据库临时变量
// 	DBS      *sql.DB             //mysql数据库全局变量
// 	REDIS_DB *redis.Client       //redis数据库全局变量
// 	Logger   *zap.Logger         //日志全局变量
// }

// var global Global

var (
	DBC    *refraction_struct.DBconf //总数据库临时变量
	DBS    *sql.DB                   //mysql数据库全局变量
	DBR    *redis.Client             //redis数据库全局变量
	Logger *zap.Logger               //日志全局变量
)
