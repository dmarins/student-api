package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	ILogger interface {
		Info(ctx context.Context, msg string, fields ...map[string]interface{})
		Error(ctx context.Context, msg string, err error, fields ...map[string]interface{})
		Fatal(ctx context.Context, msg string, err error, fields ...map[string]interface{})
		Warn(ctx context.Context, msg string, err error, fields ...map[string]interface{})
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

	zapLogger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic("failed to initialize logger")
	}

	return &Logger{
		zapLogger: zapLogger,
	}
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

func (l *Logger) Info(ctx context.Context, msg string, fields ...map[string]interface{}) {
	finalFields := make(map[string]interface{})
	if len(fields) > 0 {
		finalFields = fields[0]
	}

	l.zapLogger.Info(msg, convertToZapFields(finalFields)...)
}

func (l *Logger) Error(ctx context.Context, msg string, err error, fields ...map[string]interface{}) {
	finalFields := make(map[string]interface{})
	if len(fields) > 0 {
		finalFields = fields[0]
	}

	addErrorFields(err, finalFields)
	l.zapLogger.Error(msg, convertToZapFields(finalFields)...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, err error, fields ...map[string]interface{}) {
	finalFields := make(map[string]interface{})
	if len(fields) > 0 {
		finalFields = fields[0]
	}

	addErrorFields(err, finalFields)
	l.zapLogger.Fatal(msg, convertToZapFields(finalFields)...)
}

func (l *Logger) Warn(ctx context.Context, msg string, err error, fields ...map[string]interface{}) {
	finalFields := make(map[string]interface{})
	if len(fields) > 0 {
		finalFields = fields[0]
	}

	addErrorFields(err, finalFields)
	l.zapLogger.Warn(msg, convertToZapFields(finalFields)...)
}

func (l *Logger) Sync() error {
	return l.zapLogger.Sync()
}
