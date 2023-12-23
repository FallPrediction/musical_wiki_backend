package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"musical_wiki/global"
	"os"
)

func InitLogger() {
	logLevel, err := zapcore.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = zapcore.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), getLoggerWriter(), logLevel)
	global.Logger = zap.New(core).Sugar()
}

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
