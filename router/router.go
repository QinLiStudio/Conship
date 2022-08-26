/*
 * @Author: fzf404
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 20:40:41
 * @Description: 路由配置
 */
package router

import (
	"github.com/QinLiStudio/Conship/api"
	"github.com/QinLiStudio/Conship/config"
	"github.com/QinLiStudio/Conship/middleware"
	"github.com/QinLiStudio/Conship/pkg/logger"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	// 设置运行模式
	gin.SetMode(config.CONFIG.Mode.RunMode)

	// 初始化 Gin 引擎
	app := gin.Default()

	// 初始化 GinZap 日志
	logger.InitGinZap(app)

	// 跨域中间件
	if config.CONFIG.Cors.Enable {
		app.Use(middleware.Cors())
	}

	// 限流中间件
	if config.CONFIG.Mode.RunMode == gin.ReleaseMode {
		app.Use(middleware.LimitRoute(config.CONFIG.Limit.Request))
	}

	// 路由
	app.POST("/meta", api.Upload)
	app.PUT("/meta", api.Update)
	app.DELETE("/meta", api.Delete)
	app.GET("/meta/:path", api.Search)

	return app
}
