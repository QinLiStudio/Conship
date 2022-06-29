
//Author: lxk20021217
//Date: 2022-06-16 19:56:00
//LastEditTime: 2022-06-24 23:47:41
//LastEditors: lxk20021217
//Description: 
//FilePath: \Conship\internal\app\db\gorm.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-16 19:56:00
//LastEditTime: 2022-06-19 19:50:59
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\db\gorm.go
//是谁总是天亮了才睡

package db

import (
	"github.com/QinLiStudio/Conship/internal/app/configs"
	"gorm.io/gorm"
)


func Dsn(r configs.Config.Postgre) string {
	return "host=" + r.Host + " user=" + r.User + " password=" + r.Password + " dbname=" + r.DBName +" port=" + r.Port + " sslmode=disable TimeZone=Asia/Shanghai"
}

if db, err := gorm.Open(postgres.Open(Dsn()), &gorm.Config{}); err != nil {
	log.Fatal("连接 postgre 失败: ", err)
} else {
	db.AutoMigrate(&app.item{})
	configs.POSTGRESDB = db
	log.Print("连接 postgre 成功: ", Dsn())
}