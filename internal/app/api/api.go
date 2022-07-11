//Author: lxk20021217
//Date: 2022-06-17 19:06:58
//LastEditTime: 2022-06-21 22:40:52
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\api.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-17 19:06:58
//LastEditTime: 2022-07-11 15:12:41
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\api\api.go
//是谁总是天亮了才睡

package api

import (
	"github.com/QinLiStudio/Conship/internal/app/schema"
	"github.com/QinLiStudio/Conship/internal/app/service"
	"github.com/QinLiStudio/Conship/pkg/error"
	"github.com/QinLiStudio/Conship/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	basePath := "./file/"
	file, err := c.FormFile("content")
	if err != nil {
		error.ErrResponse(c, error.BadRequest, error.ErrBadRequest, "配置文件上传失败。", err)
		return
	}

	url := utils.RandomUrl(6)
	secret := utils.RandomSecret(6)
	filename := basePath + url
	i := schema.Item{
		Url:    url,
		Secret: secret,
	}

	service.Upload(c, file, filename, i)
}

func Delete(c *gin.Context) {
	s := c.PostForm("secret")
	service.Delete(c, s)
}

func Update(c *gin.Context) {
	s := c.PostForm("secret")
	basePath := "./file/"
	file, err := c.FormFile("content")
	if err != nil {
		error.ErrResponse(c, error.BadRequest, error.ErrBadRequest, "配置文件上传失败。", err)
		return
	}
	url := utils.RandomUrl(6)
	filename := basePath + url

	service.Update(c, s, file, filename, url)
}

func SearchUrl(c *gin.Context) {
	u := c.Param("url")

	service.SearchUrl(c, u)
}

func SearchSecret(c *gin.Context) {
	s := c.Query("secret")

	service.SearchSecret(c, s)
}
