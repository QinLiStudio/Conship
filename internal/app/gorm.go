/*
 * @Date: 2022-06-17 10:54:29
 * @LastEditors: tich425
 * @LastEditTime: 2022-06-17 19:48:45
 * @FilePath: \Conship\internal\app\gorm.go
 * @Description: 数据库启动
 */
package app

import (
	"github.com/QinLiStudio/Conship/internal/app/config"
	"github.com/QinLiStudio/Conship/pkg/gormx"
	"gorm.io/gorm"
)

/**
 * @description: 数据库初始化
 * @return {*}
 */
func InitGormDB() (*gorm.DB, func(), error) {
	cfg := config.C.Gorm
	db, err := NewGormDB()
	if err != nil {
		return nil, nil, err
	}

	cleanFunc := func() {}

	if cfg.EnableAutoMigrate {
		err = db.AutoMigrate(db)
		if err != nil {
			return nil, cleanFunc, err
		}
	}

	return db, cleanFunc, nil
}

/**
 * @description: 数据库连接实例
 * @return {*}
 */
func NewGormDB() (*gorm.DB, error) {
	cfg := config.C
	var dsn string
	dsn = config.Postgres.DSN(cfg)

	return gormx.New(&gormx.Config{
		Debug:        cfg.Gorm.Debug,
		DBType:       cfg.Gorm.DBType,
		DSN:          dsn,
		MaxIdleConns: cfg.Gorm.MaxIdleConns,
		MaxLifetime:  cfg.Gorm.MaxLifetime,
		MaxOpenConns: cfg.Gorm.MaxOpenConns,
		TablePrefix:  cfg.Gorm.TablePrefix})
}
