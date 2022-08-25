/*
 * @Author: lxk20021217
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 16:42:58
 * @Description: 应用入口
 */
package app

import (
	"github.com/QinLiStudio/Conship/internal/app/configs"
	"github.com/QinLiStudio/Conship/internal/app/db"
	"github.com/QinLiStudio/Conship/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	logger.InitZap()     // 初始化日志
	configs.InitConfig() // 初始化配置
	db.InitPostgres()    // 初始化 Postgres
	db.InitRedis()       // 初始化 Redis
	r := InitGinEngine() // 初始化 Gin 引擎
	return r
}
