//Author: lxk20021217
//Date: 2022-06-17 19:06:58
//LastEditTime: 2022-06-21 22:40:52
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\api.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-17 19:06:58
//LastEditTime: 2022-07-13 17:29:10
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\api\api.go
//是谁总是天亮了才睡

package api

import (
	"github.com/QinLiStudio/Conship/internal/app/schema"
	"github.com/QinLiStudio/Conship/internal/app/service"
	"github.com/QinLiStudio/Conship/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	p := c.PostForm("content")

	url := utils.RandomUrl(7)
	secret := utils.RandomSecret(8)
	i := schema.Item{
		Url:    url,
		Secret: secret,
	}

	service.Upload(c, p, i)
}

func Delete(c *gin.Context) {
	s := c.PostForm("secret")
	service.Delete(c, s)
}

func Update(c *gin.Context) {
	s := c.PostForm("secret")
	p := c.PostForm("content")

	service.Update(c, s, p)
}

func Search(c *gin.Context) {
	u := c.Param("path")

	service.Search(c, u)
}

