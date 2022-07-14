package app

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// 初始化 logger
func InitLogger() error {
	// 使用 lumberjack 进行日志分割归档
	lumberJackLogger := lumberjack.Logger{
		Filename:   "./conship.log", // 日志文件位置
		MaxSize:    128,             // 进行切割前日志文件的最大大小（单位：MB）
		MaxBackups: 5,               // 保留旧文件的最大个数
		MaxAge:     7,               // 保留旧文件的最大天数
		Compress:   false,           // 是否压缩旧文件
	}

	// 自定义 logger 配置
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "conship",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevelAt(zapcore.DebugLevel)

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),                                                    // 编码配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&lumberJackLogger)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 构造日志
	logger := zap.New(core)

	logger.Info("Log initialization successful")

	_, err := NewGormDB()
	if err != nil {
		return err
	}

	return nil
}
