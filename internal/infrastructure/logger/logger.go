package logger

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	ILogger interface {
		Info(ctx context.Context, msg string, fields ...dtos.Field)
		Error(ctx context.Context, msg string, err error, fields ...dtos.Field)
		Fatal(ctx context.Context, msg string, err error, fields ...dtos.Field)
		Warn(ctx context.Context, msg string, err error, fields ...dtos.Field)
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

func convertToZapFields(fields []dtos.Field) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))

	for _, field := range fields {
		zapFields = append(zapFields, zap.Any(field.Key, field.Value))
	}

	return zapFields
}

func addErrorFields(err error, fields *[]dtos.Field) {
	if err != nil {
		*fields = append(*fields, dtos.Field{Key: "error", Value: err.Error()})
		*fields = append(*fields, dtos.Field{Key: "cause", Value: err})
	}
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...dtos.Field) {
	l.zapLogger.Info(msg, convertToZapFields(fields)...)
}

func (l *Logger) Error(ctx context.Context, msg string, err error, fields ...dtos.Field) {
	finalFields := fields
	addErrorFields(err, &finalFields)
	l.zapLogger.Error(msg, convertToZapFields(finalFields)...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, err error, fields ...dtos.Field) {
	finalFields := fields
	addErrorFields(err, &finalFields)
	l.zapLogger.Fatal(msg, convertToZapFields(finalFields)...)
}

func (l *Logger) Warn(ctx context.Context, msg string, err error, fields ...dtos.Field) {
	finalFields := fields
	addErrorFields(err, &finalFields)
	l.zapLogger.Warn(msg, convertToZapFields(finalFields)...)
}

func (l *Logger) Sync() error {
	return l.zapLogger.Sync()
}
