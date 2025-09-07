package logger

import (
	"context"
	"os"

	"go.uber.org/zap"
)

var log *zap.Logger

// Init initializes the logger.
func Init() error {
	var err error
	env := os.Getenv("APP_ENV")
	if env == "prod" {
		log, err = zap.NewProduction()
	} else {
		log, err = zap.NewDevelopment()
	}
	if err != nil {
		return err
	}
	return nil
}

func Sync() {
	if log != nil {
		_ = log.Sync()
	}
}

func baseLogger(ctx context.Context) *zap.Logger {
	traceID := GetTrace(ctx)
	return log.With(
		zap.String("trace_id", traceID),
	)
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	baseLogger(ctx).Debug(msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	baseLogger(ctx).Info(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	baseLogger(ctx).Warn(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	baseLogger(ctx).Error(msg, fields...)
}
func GetLogger() *zap.Logger {
	return log
}
