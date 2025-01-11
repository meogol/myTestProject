package routes

import (
	"go.uber.org/zap"
	"meogol/pc-service/logger"
)

var routesLogger *zap.Logger

func init() {
	routesLogger = createLogger()
}

func createLogger() *zap.Logger {
	return logger.Logger.With(
		zap.String("service", "routes"),
	)
}
