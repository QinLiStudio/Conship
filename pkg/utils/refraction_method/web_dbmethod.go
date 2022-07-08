/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-07-08 21:27:58
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-07-08 22:08:03
 * @FilePath: \Conship\pkg\utils\refraction_method\web_dbmethod.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */
package refraction_method

import (
	"github.com/QinLiStudio/Conship/internal/app"
	rs "github.com/QinLiStudio/Conship/pkg/utils/refraction_struct"
)

func (as *rs.ArticleService) Search(IntID int, article rs.Article) bool {
	return app.DBS.Where("id = ?", IntID).Order("id desc").Find(&article).Error == nil
}

func (as *rs.ArticleService) SearchS(secret string, article rs.Article) bool {
	return app.DBS.Where("secret = ?", secret).Order("secret desc").Find(&article).Error == nil
}

//其余方法代写
