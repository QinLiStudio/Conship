/*
 * @Author: fzf404
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 19:29:30
 * @Description: 入口
 */

package main

import (
	"github.com/QinLiStudio/Conship/config"
	"github.com/QinLiStudio/Conship/database"
	"github.com/QinLiStudio/Conship/pkg/logger"
	"github.com/QinLiStudio/Conship/router"
)

func init() {
	logger.InitZap()        // 初始化日志
	config.InitConfig()     // 初始化配置
	database.InitPostgres() // 初始化 Postgres
	database.InitRedis()    // 初始化 Redis

}

func main() {
	app := router.InitRouter()             // 初始化路由
	app.Run(":" + config.CONFIG.Http.Port) // 启动端口监听
	logger.Info("Hello, Conship!")         // 启动成功日志
}
