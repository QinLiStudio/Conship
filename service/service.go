/*
 * @Author: fzf404
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-26 21:13:05
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

	// 校验配置文件
	if len(m.Content) > config.CONFIG.Limit.Content {
		response.Error(c, response.TooLarge, "配置文件过大")
		return
	}

	if len(m.Content) == 0 {
		response.Error(c, response.Bad, "配置文件为空")
		return
	}

	// 写入数据库
	if err := config.POSTGRESDB.Create(&m); err.Error != nil {
		response.Error(c, response.ServerErr, "配置文件写入失败")
		return
	}

	// 写入 Redis 缓存
	config.REDISDB.Set(m.Url, m.Content, time.Hour)

	response.Ok(c, response.Created, gin.H{"url": config.CONFIG.Http.Url + "/meta/" + m.Url, "secret": m.Secret}, "配置文件上传成功")
}

/**
 * @description: 删除配置文件
 * @param {*gin.Context} c gin上下文
 * @param {string} secret 配置文件secret
 * @return {*}
 */
func Delete(c *gin.Context, secret string) {
	var m meta.Meta

	if err := config.POSTGRESDB.Where("secret = ?", secret).First(&m); err.Error != nil {
		response.Error(c, response.NotFind, "配置文件未找到")
		return
	}

	if err := config.POSTGRESDB.Where("secret = ?", secret).Delete(&m); err.Error != nil {
		response.Error(c, response.ServerErr, "配置文件删除失败")
		return
	}

	// 删除 Redis 缓存
	config.REDISDB.Del(m.Url)

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

	if err := config.POSTGRESDB.Where("secret = ?", secret).First(&m); err.Error != nil {
		response.Error(c, response.NotFind, "配置文件未找到")
		return
	}

	// 校验配置文件
	if len(content) > config.CONFIG.Limit.Content {
		response.Error(c, response.TooLarge, "配置文件过大")
		return
	}

	if len(content) == 0 {
		response.Error(c, response.Bad, "配置文件为空")
		return
	}

	m.Content = content

	if err := config.POSTGRESDB.Updates(&m); err.Error != nil {
		response.Error(c, response.ServerErr, "数据库请求数据失败")
		return
	}

	// 写入 Redis 缓存
	config.REDISDB.Set(m.Url, m.Content, time.Hour)

	response.Ok(c, response.OK, gin.H{"url": config.CONFIG.Http.Url + "/meta/" + m.Url, "secret": m.Secret}, "配置文件修改成功")
}

/**
 * @description: 获取配置文件
 * @param {*gin.Context} c
 * @param {string} url
 * @return {*}
 */
func Search(c *gin.Context, url string) {

	var m meta.Meta

	// 通过 url 读取配置文件
	if len(url) == 7 {

		// 从缓存中读取
		value, err := config.REDISDB.Get(url).Result()

		if err != nil {

			// 未命中缓存
			if err := config.POSTGRESDB.Where("url = ?", url).First(&m); err.Error != nil {
				response.Error(c, response.NotFind, "配置文件未找到")
				return
			}

			// 写入 Redis 缓存
			config.REDISDB.Set(m.Url, m.Content, time.Hour)

		} else {

			// 命中缓存
			m.Content = value
		}

		c.String(200, m.Content)
		return
	}

	// 通过 secret 读取配置文件
	if len(url) == 8 {

		// 通过 secret 读取配置文件
		if err := config.POSTGRESDB.Where("secret = ?", url).First(&m); err.Error != nil {
			response.Error(c, response.NotFind, "配置文件未找到")
			return
		}

		c.String(200, m.Content)
		return
	}

	response.Error(c, response.NotFind, "配置文件未找到")
}
