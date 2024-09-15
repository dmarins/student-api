package logger

import (
	"context"
	"sync"

	"github.com/dmarins/student-api/internal/infrastructure/env"
	"go.uber.org/zap"
)

var (
	singleton *zap.Logger
	once      sync.Once
)

func NewLogger() {
	var zapLogger *zap.Logger

	if env.ProvideAppEnv() == "local" {
		zapLogger, _ = zap.NewDevelopment()
	} else {
		zapLogger, _ = zap.NewProduction()
	}

	once.Do(func() {
		singleton = zapLogger
		zap.ReplaceGlobals(singleton)
	})
}

func convertToZapFields(fields map[string]interface{}) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))

	for key, value := range fields {
		zapFields = append(zapFields, zap.Any(key, value))
	}

	return zapFields
}

func addErrorFields(err error, fields map[string]interface{}) {
	if err != nil {
		fields["error"] = err.Error()
		fields["cause"] = err
	}
}

func Info(ctx context.Context, msg string, fields ...map[string]interface{}) {
	finalFields := make(map[string]interface{})
	if len(fields) > 0 {
		finalFields = fields[0]
	}

	zap.L().Info(msg, convertToZapFields(finalFields)...)
}

func Error(ctx context.Context, msg string, err error, fields ...map[string]interface{}) {
	finalFields := make(map[string]interface{})
	if len(fields) > 0 {
		finalFields = fields[0]
	}

	addErrorFields(err, finalFields)
	zap.L().Error(msg, convertToZapFields(finalFields)...)
}

func Fatal(ctx context.Context, msg string, err error, fields ...map[string]interface{}) {
	finalFields := make(map[string]interface{})
	if len(fields) > 0 {
		finalFields = fields[0]
	}

	addErrorFields(err, finalFields)
	zap.L().Fatal(msg, convertToZapFields(finalFields)...)
}

func Warn(ctx context.Context, msg string, err error, fields ...map[string]interface{}) {
	finalFields := make(map[string]interface{})
	if len(fields) > 0 {
		finalFields = fields[0]
	}

	addErrorFields(err, finalFields)
	zap.L().Warn(msg, convertToZapFields(finalFields)...)
}

func Sync() error {
	return zap.L().Sync()
}
