package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

var logLevel = zap.NewAtomicLevel()

type Level = zapcore.Level

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
)

// 设置日志等级
func SetLevel(level Level) {
	logLevel.SetLevel(level)
}

func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return logger
	}
	if ctxLogger, ok := ctx.Value(logLevel).(*zap.Logger); ok {
		return ctxLogger
	} else {
		return logger
	}
}

func NewContext(ctx context.Context, fields ...zap.Field) context.Context {
	return context.WithValue(ctx, logLevel, WithContext(ctx).With(fields...))
}

func NewTagContext(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, "tag", tag)
}

func FromTagContext(ctx context.Context) string {
	v := ctx.Value("tag")
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
