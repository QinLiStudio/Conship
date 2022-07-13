//Author: lxk20021217
//Date: 2022-06-17 16:25:26
//LastEditTime: 2022-06-19 00:16:15
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\router.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-17 16:25:26
//LastEditTime: 2022-07-13 16:11:25
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\router\router.go
//是谁总是天亮了才睡

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
	app.Use(middleware.LimitRoute(configs.CONFIG.Limit.Limit))
	{
		app.POST("/meta", api.Upload)
		app.PUT("/meta", api.Update)
		app.DELETE("/meta", api.Delete)
		app.GET("/meta/*path", api.Search)
	}
}
