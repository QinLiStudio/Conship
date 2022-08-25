/*
 * @Author: fzf404
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 17:29:16
 * @Description: 随机数生成
 */

package utils

import (
	"math/rand"
	"time"
)

/**
 * @description: 随机链接
 * @param {int} length
 */
func RandomUrl(length int) string {
	rand.Seed(time.Now().UnixNano())
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

/**
 * @description: 随机密钥
 * @param {int} length
 */
func RandomSecret(length int) string {
	rand.Seed(time.Now().UnixNano())
	letters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
