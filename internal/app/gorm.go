/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-06-19 16:59:39
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-07-07 11:23:29
 * @FilePath: \Conship\internal\app\gorm.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */
package app

import (
	"database/sql"
	// "github.com/QinLiStudio/Conship/pkg/utils/refraction"
	_ "github.com/lib/pq"
)

//数据库变量DBS
var db *DB //临时数据库变量

db=DBS //赋值给db

var dbc *DBconf //临时解析结构体

dbc=DBC //赋值给dbc


//数据库开启链接函数
func DbOpen() {
	//sql.Open("postgres", "host=localhost port=5432 user=postgres password=123456 dbname=Test sslmode=disable")
	db, err := sql.Open(dbc.Postgres.Dsn())
	if err != nil {
		//错误信息
	} else {
		DB = db
		//提示成功
	}
}

DBS=db //还值给DBS