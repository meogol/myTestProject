package database

import (
	"go.uber.org/zap"
	"meogol/db-service/logger"
)

var dbLogger *zap.SugaredLogger

func init() {
	dbLogger = createLogger()
}

func createLogger() *zap.SugaredLogger {
	return logger.Logger.With(
		zap.String("service", "database"),
	).Sugar()
}
