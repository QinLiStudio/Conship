/*
 * @Author: git config SnowFish && git config 3200401354@qq.com
 * @Date: 2022-06-19 16:59:39
 * @LastEditors: git config SnowFish && git 3200401354@qq.com
 * @LastEditTime: 2022-06-22 19:24:06
 * @FilePath: \Conship\Conship\internal\app\logger.go
 * @Description:
 *
 * Copyright (c) 2022 by snow-fish 3200401354@qq.com, All Rights Reserved.
 */
package app

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//此Logger应在app中作为全局变量
var logger *zap.Logger

//初始化一个production型的Logger
func InitLogger() {
	logger, _ = zap.NewProduction()
}

//logger写入文件
//zapcore.Core需要三个配置——Encoder，WriteSyncer，LogLevel（写入等级）

//初始化sugarLogger
func Initfile() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	sugarLogger := logger.Sugar() //sugarLogger可作为能写入文件形式的logger变量
}

//调用sugarLogger
func Log() {
	InitLogger()

	//simpleHttpGet("www.baidu.com") 举例simpleHttpGet调用

	defer sugarLogger.Sync() //结束前写入文件
}

//写入模式
func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

//写入位置
func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./Conship/logger/test.log")
	return zapcore.AddSync(file)
}

//simpleHttpGet函数举例
/*
func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
*/
