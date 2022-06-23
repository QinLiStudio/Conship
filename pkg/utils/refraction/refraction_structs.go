/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-06-22 15:30:15
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-06-22 19:24:29
 * @FilePath: \Conship\Conship\pkg\utils\refraction\refraction_structs.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */
package refraction

//数据库映射结构体
type toDB struct {
	Host     string // 服务器地址
	Port     string // 端口
	Username string // 数据库用户名
	Password string // 数据库密码
	DBname   string // 数据库名
	SSLMode  string // SSL模式
}

//viper解析所需中间结构体
type DBconf struct {
	Postgres toDB
}
