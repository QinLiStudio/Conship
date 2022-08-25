/*
 * @Author: fzf404
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 14:40:32
 * @Description: 请求限制
 */

package middleware

import (
	"time"

	"github.com/QinLiStudio/Conship/internal/app/configs"
	"github.com/QinLiStudio/Conship/pkg/response"
	"github.com/gin-gonic/gin"
)

/**
 * @description: 次数/小时限流
 * @param {int64} limit
 * @return {*}
 */
func LimitRoute(limit int64) gin.HandlerFunc {

	return func(c *gin.Context) {

		key := c.ClientIP() + c.FullPath()

		// 读取请求次数
		value, err := configs.REDISDB.Get(key).Int64()

		if err != nil {

			// 初始化请求次数
			configs.REDISDB.Set(key, 1, time.Hour)
			c.Next()

		} else if value < limit {

			// 请求次数递增
			configs.REDISDB.Incr(key)
			c.Next()

		} else {

			// 请求次数超过限制
			response.Response(c, response.TooMany, gin.H{}, "请求过于频繁")
			c.Abort()

		}
	}
}

/**
 * @description: 单次/段时间限流
 * @param {int} limit
 */
func LimitAverageRoute(limit int64) gin.HandlerFunc {

	return func(c *gin.Context) {

		key := c.ClientIP() + c.FullPath()

		// 读取值并验证
		if err := configs.REDISDB.Get(key).Err(); err != nil {

			// 初始化请求次数
			configs.REDISDB.Set(key, 1, time.Hour/time.Duration(limit))
			c.Next()

		} else {

			response.Response(c, response.TooMany, response.TooMany, "请求过于频繁，请一小时后再重试")
			c.Abort()

		}
	}
}
