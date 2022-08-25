/*
 * @Author: fzf404
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 17:11:27
 * @Description: 服务
 */
package service

import (
	"io/ioutil"
	"os"

	"github.com/QinLiStudio/Conship/internal/app/configs"
	"github.com/QinLiStudio/Conship/internal/app/schema"
	"github.com/QinLiStudio/Conship/pkg/response"
	"github.com/gin-gonic/gin"
)

/**
 * @description: 上传配置文件
 * @param {*gin.Context} c gin上下文
 * @param {string} p 配置文件内容
 * @param {schema.Meta} i 配置文件信息
 * @return {*}
 */
func Upload(c *gin.Context, p string, i schema.Meta) {

	// 写入数据库
	if err := configs.POSTGRESDB.Select("url", "secret").Create(&i); err.Error != nil {
		response.Response(c, response.ServerErr, response.ServerErr, "数据库写入失败")
		return
	}

	// 创建文件
	file, err := os.Create("./meta/" + i.Url)
	if err != nil {
		response.Response(c, response.ServerErr, response.ServerErr, "创建文件失败")
		file.Close()
		return
	}

	// 写入文件
	if _, err := file.WriteString(p); err != nil {
		response.Response(c, response.ServerErr, response.ServerErr, "文件写入配置失败")
		file.Close()
		return
	}

	// 关闭文件
	file.Close()

	response.Response(c, response.Created, gin.H{"url": configs.CONFIG.Http.Host + "/" + i.Url, "secret": i.Secret}, "配置文件上传成功")
}

/**
 * @description: 删除配置文件
 * @param {*gin.Context} c gin上下文
 * @param {string} s 配置文件secret
 * @return {*}
 */
func Delete(c *gin.Context, s string) {
	var i schema.Meta

	if err := configs.POSTGRESDB.Table("Metas").Where("secret = ?", s).Find(&i); err.Error != nil {
		response.Response(c, response.ServerErr, response.ServerErr, "数据库请求数据失败")
		return
	}

	if err := configs.POSTGRESDB.Table("Metas").Where("secret = ?", s).Delete(&i); err.Error != nil {
		response.Response(c, response.ServerErr, response.ServerErr, "数据库删除数据失败")
		return
	}

	err := os.Remove("./meta/" + i.Url)

	if err != nil {
		response.Response(c, response.NotFind, gin.H{}, "未找到配置文件")
		return
	}

	response.Response(c, response.OK, gin.H{}, "配置文件删除成功")
}

/**
 * @description: 更新配置文件
 * @param {*gin.Context} c gin上下文
 * @param {string} s 配置文件secret
 * @param {string} p 配置文件内容
 * @return {*}
 */
func Update(c *gin.Context, s string, p string) {
	var i schema.Meta

	if err := configs.POSTGRESDB.Table("Metas").Where("secret = ?", s).Find(&i); err.Error != nil {
		response.Response(c, response.ServerErr, response.ServerErr, "数据库请求数据失败")
		return
	}

	file, err := os.OpenFile("./meta/"+i.Url, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		response.Response(c, response.NotFind, gin.H{}, "未找到配置文件")
		file.Close()
		return
	}

	if _, err := file.WriteString(p); err != nil {
		response.Response(c, response.ServerErr, response.ServerErr, "配置文件修改失败")
		file.Close()
		return
	}

	file.Close()
	response.Response(c, response.OK, gin.H{"url": configs.CONFIG.Http.Host + "/" + i.Url, "secret": i.Secret}, "配置文件修改成功")
}

/**
 * @description: 获取配置文件
 * @param {*gin.Context} c
 * @param {string} u
 * @return {*}
 */
func Search(c *gin.Context, u string) {

	if len(u) == 7 {
		file, err := os.Open("./meta/" + u)
		if err != nil {
			response.Response(c, response.NotFind, gin.H{}, "未找到配置文件")
			file.Close()
			return
		}

		content, err := ioutil.ReadAll(file)
		if err != nil {
			response.Response(c, response.ServerErr, response.ServerErr, "查询配置文件失败")
			file.Close()
			return
		}

		file.Close()
		c.String(200, string(content))
		return
	}

	if len(u) == 8 {
		var i schema.Meta
		if err := configs.POSTGRESDB.Table("Metas").Where("secret = ?", u).Find(&i); err.Error != nil {
			response.Response(c, response.ServerErr, response.ServerErr, "数据库请求数据失败")
			return
		}

		file, err := os.Open("./meta/" + i.Url)
		if err != nil {
			response.Response(c, response.NotFind, gin.H{}, "未找到配置文件")
			file.Close()
			return
		}

		content, err := ioutil.ReadAll(file)
		if err != nil {
			response.Response(c, response.NotFind, gin.H{}, "未找到配置文件")
			file.Close()
			return
		}

		file.Close()
		c.String(200, string(content))
		return
	}
	response.Response(c, response.NotFind, gin.H{}, "未找到配置文件")
}
