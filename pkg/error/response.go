//Author: lxk20021217
//Date: 2022-06-16 19:56:00
//LastEditTime: 2022-06-24 14:46:07
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\pkg\error\response.go
//是谁总是天亮了才睡


package error

import (
	"github.com/gin-gonic/gin"
)

const (
	ErrBadRequest 		= 1001 //请求错误
	ErrUnauth      		= 1002 // 未经授权
	ErrForbidden   		= 1003 // 禁止请求
	ErrNotFind     		= 1004 // 懂得都懂
	ErrTooManyRequests	= 1005 //请求频繁

	ErrServerErr 		= 0000 // 服务器错误
)
const (
	OK      		= 200 // 请求成功
	Created 		= 201 // 创建成功

	BadRequest 		= 400 //请求错误
	Unauth     		= 401 // 未经授权
	Forbidden  		= 403 // 禁止请求
	NotFind    		= 404 // 懂得都懂
	TooManyRequests = 429 //请求频繁

	ServerErr 		= 500 // 服务器错误
)

type ResponseError struct {
	Code    int    // 错误码
	Status  int    // 响应状态码
	Message string // 返回消息
	Err     error  // 响应错误
}

type ResponseNormal struct {
	Status  int         		`json:"code"`			// 响应状态码	
	Data    interface{} 		`json:"data"`			//返回内容
	Message string      		`json:"msg"`		// 返回消息
}

func ResponseNor(c *gin.Context, status int, data interface{}, message string) {
	c.JSON(status, ResponseNormal{
		Status:  status,
		Data:    data,
		Message: message,
	})
}

func ResponseErr(c *gin.Context, code int, status int, message string, err error) {
	c.JSON(status, ResponseError{
		Code:    code,
		Status:  status,
		Message: message,
		Err:     err,
	})
}

func Response(c *gin.Context, status int, data interface{}, message string) {
	ResponseNor(c, status, data, message)
}

func ErrResponse(c *gin.Context, code int, status int, message string, err error) {
	ResponseErr(c, code, status, message, err)
}
