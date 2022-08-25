/*
 * @Author: lxk20021217
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 19:11:58
 * @Description: 数据库连接
 */

package database

import (
	"github.com/QinLiStudio/Conship/config"
	"github.com/QinLiStudio/Conship/meta"
	"github.com/QinLiStudio/Conship/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() {

	if db, err := gorm.Open(postgres.Open(config.CONFIG.Postgres.Dsn()), &gorm.Config{}); err != nil {

		logger.Panic("连接 Postgres 失败: %v", err)

	} else {

		db.AutoMigrate(&meta.Meta{})

		config.POSTGRESDB = db

		logger.Info("连接 Postgres 成功: %v", config.CONFIG.Postgres)
	}
}
