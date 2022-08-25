/*
 * @Author: fzf404
 * @Date: 2022-08-18 11:01:34
 * @LastEditors: fzf404 nmdfzf404@163.com
 * @LastEditTime: 2022-08-25 17:26:00
 * @Description: Zap 初始化
 */

package logger

import (
	"os"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger
var logger *zap.Logger

// 初始化 Zap
func InitZap() {

	config := zap.NewProductionEncoderConfig()                                          // 初始化 Zap 配置
	config.EncodeTime = zapcore.RFC3339TimeEncoder                                      // 配置时间格式
	fileEncoder := zapcore.NewJSONEncoder(config)                                       // 初始化日志输出
	consoleEncoder := zapcore.NewConsoleEncoder(config)                                 // 初始化控制台输出
	logFile, _ := os.OpenFile("conship.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // 初始化日志文件
	writer := zapcore.AddSync(logFile)                                                  // 增加异步写入
	defaultLogLevel := zapcore.DebugLevel                                               // 默认日志级别
	core := zapcore.NewTee(                                                             // 日志分发通道
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)) // 初始化 Zap
	log = logger.Sugar()
}

// 初始化 GinZap 日志
func InitGinZap(r *gin.Engine) {
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
}

func Debug(template string, args ...interface{}) {
	log.Debugf(template, args...)
}

func Info(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Warn(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

func Error(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func Panic(template string, args ...interface{}) {
	log.Panicf(template, args...)
}
