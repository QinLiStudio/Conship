/*
 * @Author: fzf404
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 20:03:42
 * @Description: 响应配置
 */

package response

import (
	"github.com/gin-gonic/gin"
)

const (
	OK      = 200 // 请求成功
	Created = 201 // 创建成功

	Bad = 400 // 参数错误
	// UnAuth    = 401 // 未经授权
	// Forbidden = 403 // 禁止请求
	NotFind  = 404 // 找不到
	TooLarge = 413 // 请求过大
	TooMany  = 429 // 请求频繁

	ServerErr = 500 // 服务器错误
)

type Response struct {
	Status  int         `json:"code"` // 响应状态码
	Data    interface{} `json:"data"` //返回内容
	Message string      `json:"msg"`  // 返回消息
}

func Ok(c *gin.Context, status int, data interface{}, message string) {
	c.JSON(status, Response{
		Status:  status,
		Data:    data,
		Message: message,
	})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, Response{
		Status:  status,
		Data:    gin.H{},
		Message: message,
	})
}
