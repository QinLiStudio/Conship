
//Author: lxk20021217
//Date: 2022-06-16 19:56:00
//LastEditTime: 2022-07-02 17:32:35
//LastEditors: lxk20021217
//Description: 
//FilePath: \Conship\pkg\logger\logger.go
//是谁总是天亮了才睡

//Author: lxk20021217
//Date: 2022-06-16 19:56:00
//LastEditTime: 2022-06-24 15:14:58
//LastEditors: lxk20021217
//Description:
//FilePath: \Conship\pkg\logger\logger.go
//是谁总是天亮了才睡

package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger

type Level = zapcore.Level

const (
	DebugLevel Level = -1

	InfoLevel Level = iota
	WarnLevel
	ErrorLevel
	PanicLevel
)

func InitZap() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	highlevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= zap.ErrorLevel
	})
	lowlevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l < zap.ErrorLevel && l >= zap.DebugLevel
	})

	lowFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./tem/low.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	})
	lowFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(lowFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowlevel)

	highFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./tem/high.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	})
	highFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(highFileWriteSyncer, zapcore.AddSync(os.Stdout)), highlevel)

	core := zapcore.NewTee(highFileCore, lowFileCore)
	logger := zap.New(core, zap.AddCaller())
	log = logger.Sugar()
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
