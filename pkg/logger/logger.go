/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-06-19 16:59:39
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-07-07 11:45:13
 * @FilePath: \Conship\pkg\logger\logger.go
 * @Description: 
 * 
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved. 
 */
package logger

import (
	"go.uber.org/zap"
)

//应来自app的全局logger
var logger2 *zap.Logger

logger2=Logger

//func (log *Logger) MethodXXX(msg string, fields ...Field)    日志记录函数语法
/*
func (log *Logger) Error(msg string, fields ...Field)
func (log *Logger) Info(msg string, fields ...Field)
*/

//zapcore.Field的别名是Field  每个zapcore.Field其实就是一组键值对参数
// type Field = zapcore.Field

//关于logger的相应事例

  Logger2.Debug("this is debug message")
  logger2.Info("this is info message")
  logger2.Info("this is info message with fileds",
	zap.Int("code", 400), zap.String("error", "..."))
  logger2.Warn("this is warn message")
  logger2.Error("this is error message")
  logger2.Panic("this is panic message")



