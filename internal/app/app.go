//Author: lxk20021217
//Date: 2022-06-16 19:56:00
//LastEditTime: 2022-07-08 15:08:30
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\app.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-16 19:56:00
//LastEditTime: 2022-07-09 16:19:32
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\app.go
//是谁总是天亮了才睡

package app

import (
	"github.com/QinLiStudio/Conship/internal/app/configs"
	"github.com/QinLiStudio/Conship/internal/app/db"
	"github.com/QinLiStudio/Conship/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	logger.InitZap()
	configs.InitConfig()
	db.InitPostgres()
	db.InitRedis()
	r := InitGinEngine()
	return r
}
