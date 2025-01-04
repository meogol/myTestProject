package config

import (
	"go.uber.org/zap"
	"meogol/db-service/logger"
)

var configLogger *zap.SugaredLogger

func init() {
	configLogger = createLogger()
}

func createLogger() *zap.SugaredLogger {
	return logger.Logger.With(
		zap.String("service", "config"),
	).Sugar()
}
