
//Author: lxk20021217
//Date: 2022-06-16 19:56:00
//LastEditTime: 2022-06-19 00:44:30
//LastEditors: lxk20021217
//Description: 
//FilePath: \Conship\pkg\res\response.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-16 19:56:00
//LastEditTime: 2022-06-17 15:48:06
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\pkg\errors\response.go
//是谁总是天亮了才睡

package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	Success = 200 // 成功

	NoLogin         = 401 // 未登录
	ParamErr        = 402 // 参数错误
	NoAuth          = 403 // 无权限
	NotFind         = 404 // 找不到
	TokenExpired    = 405 // Token过期
	TooManyRequests = 429 // 请求过于频繁

	ServerErr = 500 // 服务器错误
	TokenErr  = 501 // Token分发失败
	SQLErr    = 502 // sql错误
)

/**
 * @description: 响应结果
 * @param {int} code
 * @param {interface{}} data
 * @param {string} msg
 * @param {*gin.Context} c
 */
func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(Success, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(Success, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(Success, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(Success, data, message, c)
}

func Fail(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(code int, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}

func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}
