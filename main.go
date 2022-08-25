/*
 * @Author: lxk20021217
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-18 18:05:51
 * @Description: 入口函数
 */

package main

import (
	"github.com/QinLiStudio/Conship/internal/app"
	"github.com/QinLiStudio/Conship/internal/app/configs"
	"github.com/QinLiStudio/Conship/pkg/logger"
)

func main() {
	r := app.Init()                       // 初始化
	r.Run(":" + configs.CONFIG.Http.Port) // 启动端口监听
	logger.Info("Hello, Conship!")        // 启动成功日志
}
