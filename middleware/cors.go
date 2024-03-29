/*
 * @Author: fzf404
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 14:11:52
 * @Description: 跨域配置
 */

package middleware

import (
	"github.com/QinLiStudio/Conship/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {

	c := cors.DefaultConfig()

	c.AllowMethods = config.CONFIG.Cors.AllowMethods
	c.AllowHeaders = config.CONFIG.Cors.AllowHeaders
	c.AllowOrigins = config.CONFIG.Cors.AllowOrigins

	return cors.New(c)
}
