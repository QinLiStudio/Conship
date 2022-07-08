/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-07-08 20:48:23
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-07-08 22:05:59
 * @FilePath: \Conship\pkg\utils\refraction_method\web_method_middle.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */
package refraction_method

import (
	rs "github.com/QinLiStudio/Conship/pkg/utils/refraction_struct"
	"github.com/gin-gonic/gin"
)

func SearchFrom_ID(c *gin.Context) {

	ID := c.Query("id") // 读取 ID 参数

	rs.ArticleService.IDSearch(ID, c)
}

func SearchFrom_Secret(c *gin.Context) {

	ID := c.Query("secret") // 读取 secret 参数

	rs.ArticleService.SecretSearch(secret, c)
}

func UpLoad(c *gin.Context) {

	var a rs.Article

	if err := c.ShouldBind(&a); err != nil {

		response.FailWithDetailed(response.ParamErr, err.Error(), "提交信息非法", c)

		return
	}

	rs.ArticleService.ToUpLoad(a, c)
}

func Refresh(c *gin.Context) {

	var a rs.Article

	if err := c.ShouldBind(&a); err != nil {

		response.FailWithDetailed(response.ParamErr, err.Error(), "提交信息非法", c)

		return
	}

	rs.ArticleService.ToRefresh(a, c)
}

func Delete(c *gin.Context) {

	var a rs.Article

	if err := c.ShouldBind(&a); err != nil {

		response.FailWithDetailed(response.ParamErr, err.Error(), "提交信息非法", c)

		return
	}

	rs.ArticleService.ToDelete(a, c)
}
