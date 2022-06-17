/*
 * @Date: 2022-06-17 17:15:07
 * @LastEditors: tich425
 * @LastEditTime: 2022-06-17 19:46:46
 * @FilePath: \Conship\internal\app\db\db.go
 * @Description: 数据库自动建表功能
 */
package db

import (
	"strings"

	"github.com/QinLiStudio/Conship/internal/app/config"
	"github.com/QinLiStudio/Conship/internal/app/entity"
	"gorm.io/gorm"
)

/**
 * @description: gorm 数据库自动建表
 * @param {*gorm.DB} db
 * @return {*}
 */
func AutoMigrate(db *gorm.DB) error {
	if dbType := config.C.Gorm.DBType; strings.ToLower(dbType) == "postgres" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}

	return db.AutoMigrate(new(entity.Conship))
}
