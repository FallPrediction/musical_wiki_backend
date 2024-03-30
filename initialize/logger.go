package initialize

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger
var loggerOnce sync.Once

func getLoggerWriter() zapcore.WriteSyncer {
	os.MkdirAll("./logs", os.ModePerm)
	file, _ := os.Create("./logs/gin.log")
	return zapcore.AddSync(file)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func NewLogger() *zap.SugaredLogger {
	if logger == nil {
		loggerOnce.Do(func() {
			logLevel, err := zapcore.ParseLevel(os.Getenv("LOG_LEVEL"))
			if err != nil {
				logLevel = zapcore.InfoLevel
			}
			core := zapcore.NewCore(getEncoder(), getLoggerWriter(), logLevel)
			logger = zap.New(core).Sugar()
		})
	}
	return logger
}
