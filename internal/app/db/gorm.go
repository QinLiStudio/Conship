/*
 * @Author: lxk20021217
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 16:57:44
 * @Description: 数据库连接
 */

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

		logger.Panic("连接 Postgres 失败: %v", err)

	} else {

		db.AutoMigrate(&schema.Meta{})

		configs.POSTGRESDB = db

		logger.Info("连接 Postgres 成功: %v", configs.CONFIG.Postgres)
	}
}
