/*
 * @Date: 2022-06-17 17:26:25
 * @LastEditors: tich425
 * @LastEditTime: 2022-06-17 19:50:53
 * @FilePath: \Conship\pkg\gormx\gormx.go
 * @Description:
 */
package gormx

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/**
 * @description: gorm 配置结构体
 * @return {*}
 */
type Config struct {
	Debug        bool
	DBType       string
	DSN          string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	TablePrefix  string
}

/**
 * @description: 创建 gorm 实例
 * @param {*Config} c
 * @return {*}
 */
func New(c *Config) (*gorm.DB, error) {
	var dialector gorm.Dialector
	dialector = postgres.Open(c.DSN)
	gconfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.TablePrefix,
			SingularTable: true,
		},
	}

	db, err := gorm.Open(dialector, gconfig)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 写入 toml 中的访问限制
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)

	return db, nil
}
