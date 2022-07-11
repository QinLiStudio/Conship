//Author: lxk20021217
//Date: 2022-06-16 19:56:00
//LastEditTime: 2022-07-09 14:52:51
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\db\gorm.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-16 19:56:00
//LastEditTime: 2022-07-10 15:50:11
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\db\gorm.go
//是谁总是天亮了才睡

package db

import (
	"github.com/QinLiStudio/Conship/internal/app/configs"
	"github.com/QinLiStudio/Conship/internal/app/schema"
	"github.com/QinLiStudio/Conship/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() {
	r := configs.CONFIG.Postgres
	if db, err := gorm.Open(postgres.Open(configs.Dsn(r)), &gorm.Config{}); err != nil {
		logger.Error("连接 Postgres 失败: %v", err)
	} else {
		db.AutoMigrate(&schema.Item{})
		configs.POSTGRESDB = db
		logger.Info("连接 Postgres 成功: %v", configs.Dsn(r))
	}
}
