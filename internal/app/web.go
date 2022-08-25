/*
 * @Author: lxk20021217
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 17:05:15
 * @Description: 应用入口
 */

package app

import (
	"github.com/QinLiStudio/Conship/internal/app/configs"

	"github.com/QinLiStudio/Conship/internal/app/middleware"
	"github.com/QinLiStudio/Conship/internal/app/router"
	"github.com/QinLiStudio/Conship/pkg/logger"
	"github.com/gin-gonic/gin"
)

func InitGinEngine() *gin.Engine {

	// 设置运行模式
	gin.SetMode(configs.CONFIG.Mode.RunMode)

	// 初始化 Gin 引擎
	app := gin.Default()

	// 初始化 GinZap 日志
	logger.InitGinZap(app)
	
	// 跨域中间件
	if configs.CONFIG.Cors.Enable {
		app.Use(middleware.Cors())
	}
	// 注册路由
	router.Register(app)

	return app
}
