//Author: lxk20021217
//Date: 2022-06-16 19:56:00
//LastEditTime: 2022-07-09 15:46:23
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\cmd\main.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-16 19:56:00
//LastEditTime: 2022-07-09 16:27:59
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\cmd\main.go
//是谁总是天亮了才睡

package main

import (
	"github.com/QinLiStudio/Conship/internal/app"
	"github.com/QinLiStudio/Conship/internal/app/configs"
	"github.com/QinLiStudio/Conship/pkg/logger"
)

func main() {
	r := app.Init()
	r.Run(":" + configs.CONFIG.Http.Port) // 启动服务
	logger.Info("Hello,Conship!")
}
