/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-06-22 15:46:44
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-07-08 18:27:30
 * @FilePath: \Conship\pkg\utils\refraction_method\refraction_method.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */
package refraction_method

import (
	"log"
	"github.com/spf13/viper"
	"github.com/QinLiStudio/Conship/internal/app"
)

//DSN配置mysql返回String
func (todb *ToDB) Dsn() string {
	return "\"" + todb.Username + "\"" + "," + "\"" + "host=" + todb.Host + " " + "port=" + todb.Port + " " + "user=" + todb.Username + " " + "password=" + todb.Password + " " + "dbname=" + todb.DBname + "sslmoded=" + todb.SSLMode + "\""
}

//Viper解析方法

var Postgres *DBconf //解析所需类型结构体

Postgres=app.DBC //赋值

func Config() {
	// 文件路径
	viper.AddConfigPath("./Conship/config")
	// 文件信息
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	// 读取配置信息
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Read config failed: ", err)
	}
	// 解析至结构体
	if err := viper.Unmarshal(&Postgres); err != nil {
		log.Fatal("Unmarshal config failed: ", err)
	}
}

 Postgres.Postgres0=app.DBS //还值mysql
 Postgres.Postgres1=app.DBR //还值redis
