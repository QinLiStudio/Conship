/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-06-19 16:59:39
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-07-08 22:17:30
 * @FilePath: \Conship\internal\app\web.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */
package app

import (
	"github.com/gin-gonic/gin" 
	// "github.com/QinLiStudio/Conship/pkg/utils/refraction_method" 循环引包 寄
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	route := r.Group("/meta"){
		route.GET("/", api.Regist)          //ID搜索
		route.POST("/", api.Login)           //Secret是普索
		route.POST("/", api.SendEmailCode)    //上传
		route.POST("/", api.SendEmailCode2)  //更新
		route.POST("/", api.SendEmailCode2)  //删除
	}
	
}
