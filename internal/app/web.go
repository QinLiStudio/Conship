//Author: lxk20021217
//Date: 2022-06-21 22:04:30
//LastEditTime: 2022-07-05 15:08:16
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\web.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-21 22:04:30
//LastEditTime: 2022-07-08 21:31:57
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\web.go
//是谁总是天亮了才睡

package app

import (
	"github.com/QinLiStudio/Conship/internal/app/configs"
	"github.com/QinLiStudio/Conship/internal/app/middleware"
	"github.com/QinLiStudio/Conship/internal/app/router"
	"github.com/gin-gonic/gin"
)

func InitGinEngine() *gin.Engine {
	gin.SetMode(configs.CONFIG.Mode.Runmode)

	app := gin.Default()

	// CORS
	if configs.CONFIG.Cors.Enable {
		app.Use(middleware.Cors())
	}

	// Router register
	router.Register(app)

	return app
}
