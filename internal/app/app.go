/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-06-19 16:59:39
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-07-07 11:29:16
 * @FilePath: \Conship\internal\app\app.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */
package app

import (
	"database/sql"
	// "github.com/QinLiStudio/Conship/pkg/utils/refraction"
	"go.uber.org/zap"
)

var (
	DBC    *DBconf     //解析临时变量
	DBS    *sql.DB     //数据库全局变量
	Logger *zap.Logger //日志全局变量

)
