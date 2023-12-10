package lib

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Log() *zap.Logger {

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	logFile, _ := os.OpenFile("log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
	)
	return zap.New(core)
}
