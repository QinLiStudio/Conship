//Author: lxk20021217
//Date: 2022-06-17 16:25:26
//LastEditTime: 2022-06-19 00:16:15
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\router.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-17 16:25:26
//LastEditTime: 2022-06-19 00:45:38
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\router.go
//是谁总是天亮了才睡

package app

import (
	"github.com/QinLiStudio/Conship/internal/app"
	"github.com/QinLiStudio/Conship/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	r.POST("/meta", app.Upload)
	r.DELETE("/meta", app.Delete)
	r.PUT("/meta", app.Update)
	r.GET("/meta", app.SearchID)
	r.GET("/meta/*id", app.SearchSecret)

	return r
}
