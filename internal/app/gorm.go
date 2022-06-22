/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-06-19 16:59:39
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-06-22 19:23:57
 * @FilePath: \Conship\Conship\internal\app\gorm.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */
package app

import (
	"database/sql"

	_ "github.com/lib/pq"
	//引入包refraction存在bug无法引入
)

//数据库全局变量DB 应在app中的
var DB *sql.DB

//无法引入包变量无法识别（该todb应在app作为全局变量）
var todb toDB

//从app中引入调用Config()后的todo全局变量

//数据库开启链接函数
func DbOpen() {
	//sql.Open("postgres", "host=localhost port=5432 user=postgres password=123456 dbname=Test sslmode=disable")
	db, err := sql.Open(todb.Dsn())
	if err != nil {
		//错误信息
	} else {
		DB = db
		//提示成功
	}
}
