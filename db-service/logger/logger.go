package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

func init() {
	// logger := zap.Must(zap.NewDevelopment())
	// if os.Getenv("APP_ENV") == "production" {
	// 	logger = zap.Must(zap.NewProduction())
	// }
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	if os.Getenv("APP_ENV") == "production" {
		config = zap.NewProductionConfig()
	}
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	Logger = logger
}
