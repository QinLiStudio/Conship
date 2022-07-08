/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-06-22 15:30:15
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-07-08 18:26:13
 * @FilePath: \Conship\pkg\utils\refraction_struct\refraction_structs.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */

package refraction_struct

//数据库映射结构体
type ToDB struct {
	Host     string // 服务器地址
	Port     string // 端口
	Username string // 数据库用户名
	Password string // 数据库密码
	DBname   string // 数据库名
	SSLMode  string // SSL模式
}

type ToRedis struct {
	Addr     string // 服务器地址
	Password string // 数据库密码
	DB       int    // 数据库
}

//viper解析所需中间结构体
type DBconf struct {
	Postgres0 ToDB
	Postgres1 ToRedis
}