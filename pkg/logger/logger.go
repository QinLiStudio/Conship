/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-06-19 16:59:39
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-07-07 11:39:52
 * @FilePath: \Conship\internal\app\logger.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */
package loger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//局部变量logger1
var logger1 *zap.Logger

//初始化一个production型的Logger
func InitLogger() {
	logger1, _ = zap.NewProduction()
}

//logger写入文件
//zapcore.Core需要三个配置——Encoder，WriteSyncer，LogLevel（写入等级）

//写入模式
func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

//写入位置
func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./Conship/logger/test.log")
	return zapcore.AddSync(file)
}

//初始化Logger1
func Initfile() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger1 := zap.New(core)
}

//调用生成Logger1
func Log() {
	InitLogger()
}

Logger=logger1 //还值给Logger
