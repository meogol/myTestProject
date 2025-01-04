package database

import (
	"go.uber.org/zap"
	"meogol/db-service/logger"
)

var backupsLogger *zap.SugaredLogger

func init() {
	backupsLogger = createLogger()
}

func createLogger() *zap.SugaredLogger {
	return logger.Logger.With(
		zap.String("service", "database"),
	).Sugar()
}
