/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-06-19 16:59:39
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-06-22 19:24:14
 * @FilePath: \Conship\Conship\pkg\logger\logger.go
 * @Description: 
 * 
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved. 
 */
package logger

import (
	"go.uber.org/zap"
)

//应来自app的全局logger
var logger *zap.Logger
logger, _ = zap.NewProduction()


//func (log *Logger) MethodXXX(msg string, fields ...Field)    日志记录函数语法
/*
func (log *Logger) Error(msg string, fields ...Field)
func (log *Logger) Info(msg string, fields ...Field)
*/

//zapcore.Field的别名是Field  每个zapcore.Field其实就是一组键值对参数
// type Field = zapcore.Field

//关于logger的相应事例

  Logger.Debug("this is debug message")
  logger.Info("this is info message")
  logger.Info("this is info message with fileds",
	zap.Int("code", 400), zap.String("error", "..."))
  logger.Warn("this is warn message")
  logger.Error("this is error message")
  logger.Panic("this is panic message")



