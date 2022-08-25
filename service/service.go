/*
 * @Author: fzf404
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 20:46:35
 * @Description: 服务
 */
package service

import (
	"time"

	"github.com/QinLiStudio/Conship/config"
	"github.com/QinLiStudio/Conship/meta"
	"github.com/QinLiStudio/Conship/pkg/response"
	"github.com/gin-gonic/gin"
)

/**
 * @description: 上传配置文件
 * @param {*gin.Context} c gin上下文
 * @param {string} p 配置文件内容
 * @param {schema.Meta} m 配置文件信息
 * @return {*}
 */
func Upload(c *gin.Context, m meta.Meta) {

	// 写入数据库
	if err := config.POSTGRESDB.Create(&m); err.Error != nil {
		response.Error(c, response.ServerErr, "数据库写入失败")
		return
	}
	// 写入 Redis
	config.REDISDB.Set(m.Secret, m.Content, time.Hour)

	response.Ok(c, response.Created, gin.H{"url": config.CONFIG.Http.Host + "/" + m.Url, "secret": m.Secret}, "配置文件上传成功")
}

/**
 * @description: 删除配置文件
 * @param {*gin.Context} c gin上下文
 * @param {string} secret 配置文件secret
 * @return {*}
 */
func Delete(c *gin.Context, secret string) {
	var m meta.Meta

	if err := config.POSTGRESDB.Where("secret = ?", secret).Find(&m); err.Error != nil {
		response.Error(c, response.ServerErr, "未找到配置文件")
		return
	}

	if err := config.POSTGRESDB.Where("secret = ?", secret).Delete(&m); err.Error != nil {
		response.Error(c, response.ServerErr, "删除数据失败")
		return
	}

	response.Ok(c, response.OK, gin.H{}, "配置文件删除成功")
}

/**
 * @description: 更新配置文件
 * @param {*gin.Context} c gin上下文
 * @param {string} secret 配置文件secret
 * @param {string} p 配置文件内容
 * @return {*}
 */
func Update(c *gin.Context, secret string, content string) {
	var m meta.Meta

	if err := config.POSTGRESDB.Where("secret = ?", secret).Find(&m); err.Error != nil {
		response.Ok(c, response.ServerErr, response.ServerErr, "未找到配置文件")
		return
	}
	m.Content = content
	if err := config.POSTGRESDB.Updates(&m); err.Error != nil {
		response.Ok(c, response.ServerErr, response.ServerErr, "未找到配置文件")
		return
	}

	response.Ok(c, response.OK, gin.H{"url": config.CONFIG.Http.Host + "/" + m.Url, "secret": m.Secret}, "配置文件修改成功")
}

/**
 * @description: 获取配置文件
 * @param {*gin.Context} c
 * @param {string} url
 * @return {*}
 */
func Search(c *gin.Context, url string) {

	if len(url) == 7 {
		var m meta.Meta
		if err := config.POSTGRESDB.Where("url = ?", url).Find(&m); err.Error != nil {
			response.Error(c, response.ServerErr, "数据库请求数据失败")
			return
		}

		c.String(200, m.Content)
		return
	}

	if len(url) == 8 {
		var m meta.Meta
		if err := config.POSTGRESDB.Where("secret = ?", url).Find(&m); err.Error != nil {
			response.Ok(c, response.ServerErr, response.ServerErr, "数据库请求数据失败")
			return
		}

		c.String(200, m.Content)
		return
	}
	response.Ok(c, response.NotFind, gin.H{}, "未找到配置文件")
}
