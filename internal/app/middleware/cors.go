//Author: lxk20021217
//Date: 2022-06-17 16:27:53
//LastEditTime: 2022-06-19 19:33:20
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\internal\app\middleware\cors.go
//是谁总是天亮了才睡

package middleware

import (
	"regexp"

	"github.com/QinLiStudio/Conship/internal/app/configs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	c := cors.DefaultConfig()
	c.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	c.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	if gin.Mode() == gin.ReleaseMode {
		// 生产环境配置跨域域名，否则 403
		c.AllowOrigins = []string{configs.CONFIG.Http.Host}
	} else {
		// 测试环境允许域名的请求
		c.AllowOrigins = []string{configs.CONFIG.Http.Host}
		// 测试环境下允许 localhost 的请求
		c.AllowOriginFunc = func(origin string) bool {
			if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
				return true
			}
			if regexp.MustCompile(`^http://localhost:\d+$`).MatchString(origin) {
				return true
			}
			return false
		}
	}
	return cors.New(c)
}
