//Author: lxk20021217
//Date: 2022-06-17 16:25:26
//LastEditTime: 2022-06-19 00:16:15
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\router.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-17 16:25:26
//LastEditTime: 2022-06-30 00:19:21
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\router.go
//是谁总是天亮了才睡

package app

import (
	"github.com/QinLiStudio/Conship/internal/app/api"
	"github.com/QinLiStudio/Conship/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	r.POST("/meta", api.Upload)
	r.DELETE("/meta", api.Delete)
	r.PUT("/meta", api.Update)
	r.GET("/meta", api.SearchUrl)
	r.GET("/meta/*id", api.SearchSecret)

	return r
}
