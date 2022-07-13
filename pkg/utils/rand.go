//Author: lxk20021217
//Date: 2022-06-20 00:32:05
//LastEditTime: 2022-06-24 16:19:02
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\pkg\utils\rand.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-20 00:32:05
//LastEditTime: 2022-07-13 21:09:25
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\pkg\utils\rand.go
//是谁总是天亮了才睡

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
