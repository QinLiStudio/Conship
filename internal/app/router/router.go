/*
 * @Author: lxk20021217
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 15:37:57
 * @Description: 路由配置
 */
package router

import (
	"github.com/QinLiStudio/Conship/internal/app/api"
	"github.com/QinLiStudio/Conship/internal/app/configs"
	"github.com/QinLiStudio/Conship/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func Register(app *gin.Engine) error {

	RegisterAPI(app)
	return nil

}

func RegisterAPI(app *gin.Engine) {
	// 限流
	app.Use(middleware.LimitRoute(configs.CONFIG.Limit.Limit))
	{
		app.POST("/meta", api.Upload)
		app.PUT("/meta", api.Update)
		app.DELETE("/meta", api.Delete)
		app.GET("/meta/:path", api.Search)
	}
}
