package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var structLogger *StructLogger = NewLogger()

type Logger interface {
	Error(msg string, params ...interface{})
}

type StructLogger struct {
	logger *zap.Logger
}

type loggerField struct {
	field zapcore.Field
}

func (z *StructLogger) Error(msg string, params ...loggerField) {
	var fields []zapcore.Field
	for _, param := range params {
		fields = append(fields, param.field)
	}
	z.logger.Error(msg, fields...)
}

func (z *StructLogger) Info(msg string, params ...loggerField) {
	var fields []zapcore.Field
	for _, param := range params {
		fields = append(fields, param.field)
	}
	z.logger.Info(msg, fields...)
}

func NewLogger() *StructLogger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	return &StructLogger{
		logger: zap.Must(cfg.Build()),
	}
}

func GetLogger() StructLogger {
	return *structLogger
}

func String(key, value string) loggerField {
	return loggerField{
		field: zapcore.Field{
			Key:    key,
			Type:   zapcore.StringType,
			String: value,
		},
	}
}

func Int(key string, value int) loggerField {
	return loggerField{
		field: zapcore.Field{
			Key:     key,
			Type:    zapcore.Int64Type,
			Integer: int64(value),
		},
	}
}
