package pc

import (
	"go.uber.org/zap"
	"meogol/pc-service/logger"
)

var pcLogger *zap.SugaredLogger

func init() {
	pcLogger = createLogger()
}

func createLogger() *zap.SugaredLogger {
	return logger.Logger.With(
		zap.String("service", "pc"),
	).Sugar()
}
