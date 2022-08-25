/*
 * @Author: lxk20021217
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 19:14:24
 * @Description: 请求处理
 */
package api

import (
	"github.com/QinLiStudio/Conship/meta"
	"github.com/QinLiStudio/Conship/pkg/utils"
	"github.com/QinLiStudio/Conship/service"
	"github.com/gin-gonic/gin"
)

// 上传配置文件
func Upload(c *gin.Context) {

	// 文件内容
	content := c.PostForm("content")

	// 生成链接
	url := utils.RandomUrl(7)

	// 生成密钥
	secret := utils.RandomSecret(8)

	// 数据库信息
	m := meta.Meta{
		Url:     url,
		Secret:  secret,
		Content: content,
	}

	service.Upload(c, m)
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

	url := c.Param("path")

	service.Search(c, url)
}
