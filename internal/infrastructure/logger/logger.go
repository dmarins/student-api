package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	ILogger interface {
		Info(ctx context.Context, msg string, fields ...string)
		Error(ctx context.Context, msg string, err error, fields ...string)
		Fatal(ctx context.Context, msg string, err error, fields ...string)
		Warn(ctx context.Context, msg string, fields ...string)
		Sync() error
	}

	Logger struct {
		zapLogger *zap.Logger
	}
)

func NewLogger() ILogger {
	config := zap.NewProductionConfig()

	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	config.Encoding = "console"
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}

	zapLogger, _ := config.Build(zap.AddCallerSkip(1))

	return &Logger{
		zapLogger: zapLogger,
	}
}

func convertStringFields(extraFields []string) []zap.Field {
	if len(extraFields)%2 != 0 {
		return nil
	}

	zapFields := make([]zap.Field, 0, len(extraFields)/2)
	for i := 0; i < len(extraFields); i += 2 {
		zapFields = append(zapFields, zap.String(extraFields[i], extraFields[i+1]))
	}

	return zapFields
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...string) {
	zapFields := convertStringFields(fields)

	l.zapLogger.Info(msg, zapFields...)
}

func (l *Logger) Error(ctx context.Context, msg string, err error, fields ...string) {
	zapFields := convertStringFields(fields)
	zapFields = append(zapFields, zap.Error(err))

	l.zapLogger.Error(msg, zapFields...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, err error, fields ...string) {
	zapFields := convertStringFields(fields)
	zapFields = append(zapFields, zap.Error(err))

	l.zapLogger.Fatal(msg, zapFields...)
}

func (l *Logger) Warn(ctx context.Context, msg string, fields ...string) {
	zapFields := convertStringFields(fields)

	l.zapLogger.Warn(msg, zapFields...)
}

func (l *Logger) Sync() error {
	return l.zapLogger.Sync()
}
