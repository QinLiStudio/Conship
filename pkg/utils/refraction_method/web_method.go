/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-07-08 20:50:55
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-07-08 22:06:40
 * @FilePath: \Conship\pkg\utils\refraction_method\web_method.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */
package refraction_method

import (
	"strconv"

	rs "github.com/QinLiStudio/Conship/pkg/utils/refraction_struct"
	"github.com/gin-gonic/gin"
)

func (as *rs.ArticleService) IDSearch(ID string, c *gin.Context) {

	var article rs.Article

	IntID, err := strconv.Atoi(ID)

	if err != nil {
		//错误
	} else {
		if ok, article := as.Search(IntID, article); !ok {
			//失败

		} else {
			if len(article) == 0 {
				//失败
				return
			}
		}

		//返回json
		response.OkWithData(gin.H{
			"ID":      article.ID,
			"Content": article.Content,
		}, c)
	}
}

//其余方法待写
