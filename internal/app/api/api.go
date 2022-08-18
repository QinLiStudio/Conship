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
