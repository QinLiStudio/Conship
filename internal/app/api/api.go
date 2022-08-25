/*
 * @Author: lxk20021217
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 15:23:56
 * @Description: 请求处理
 */
package api

import (
	"github.com/QinLiStudio/Conship/internal/app/schema"
	"github.com/QinLiStudio/Conship/internal/app/service"
	"github.com/QinLiStudio/Conship/pkg/utils"
	"github.com/gin-gonic/gin"
)

// 上传配置文件
func Upload(c *gin.Context) {

	// 文件内容
	p := c.PostForm("content")

	// 生成链接
	url := utils.RandomUrl(7)

	// 生成密钥
	secret := utils.RandomSecret(8)

	// 数据库信息
	i := schema.Meta{
		Url:    url,
		Secret: secret,
	}

	service.Upload(c, p, i)
}

// 删除配置文件
func Delete(c *gin.Context) {

	s := c.PostForm("secret")

	service.Delete(c, s)

}

// 更新配置文件
func Update(c *gin.Context) {

	s := c.PostForm("secret")
	p := c.PostForm("content")

	service.Update(c, s, p)
}

// 查询配置文件
func Search(c *gin.Context) {

	u := c.Param("path")

	service.Search(c, u)
}
