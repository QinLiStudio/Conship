/*
 * @Author: fzf404
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 15:32:55
 * @Description: 响应配置
 */

package response

import (
	"github.com/gin-gonic/gin"
)

const (
	OK      = 200 // 请求成功
	Created = 201 // 创建成功

	// Bad       = 400 // 请求错误
	// UnAuth    = 401 // 未经授权
	// Forbidden = 403 // 禁止请求
	NotFind = 404 // 懂得都懂
	TooMany = 429 // 请求频繁

	ServerErr = 500 // 服务器错误
)

type ResponseNormal struct {
	Status  int         `json:"code"` // 响应状态码
	Data    interface{} `json:"data"` //返回内容
	Message string      `json:"msg"`  // 返回消息
}

func Response(c *gin.Context, status int, data interface{}, message string) {
	c.JSON(status, ResponseNormal{
		Status:  status,
		Data:    data,
		Message: message,
	})
}
