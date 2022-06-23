/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-06-19 16:59:39
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-06-22 19:23:44
 * @FilePath: \Conship\Conship\internal\app\app.go
 * @Description: 
 * 
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved. 
 */
package app

var (
	var todb toDB //数据库全局变量（未配置前身）
	var DB *sql.DB //数据库全局变量（配置后）
	var logger *zap.Logger //日志全局变量

)