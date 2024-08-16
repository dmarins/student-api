package logger

import (
	"go.uber.org/zap"
)

func InitLogger() (*zap.Logger, error) {
	Logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return Logger, nil
}
