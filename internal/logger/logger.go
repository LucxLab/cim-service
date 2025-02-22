package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	DebugWithData(message string, data map[string]interface{})
	InfoWithData(message string, data map[string]interface{})
	WarnWithData(message string, data map[string]interface{})
	ErrorWithData(message string, data map[string]interface{})
	Close() error
}

type zapLogger struct {
	efficient *zap.Logger
	sugared   *zap.SugaredLogger
}

func (l *zapLogger) DebugWithData(message string, data map[string]interface{}) {
	l.logWithData(zapcore.DebugLevel, message, data)
}

func (l *zapLogger) InfoWithData(message string, data map[string]interface{}) {
	l.logWithData(zapcore.InfoLevel, message, data)
}

func (l *zapLogger) WarnWithData(message string, data map[string]interface{}) {
	l.logWithData(zapcore.WarnLevel, message, data)
}

func (l *zapLogger) ErrorWithData(message string, data map[string]interface{}) {
	l.logWithData(zapcore.ErrorLevel, message, data)
}

func (l *zapLogger) logWithData(level zapcore.Level, message string, data map[string]interface{}) {
	keysAndValues := make([]interface{}, 0, len(data)*2)
	for key, value := range data {
		keysAndValues = append(keysAndValues, key, value)
	}
	l.sugared.Logw(level, message, keysAndValues...)
}

func (l *zapLogger) Close() error {
	return l.efficient.Sync()
}

func New(isDevelopment bool) (Logger, error) {
	var logger *zap.Logger
	var loggerErr error
	var utcClock = newUtcClock()

	if isDevelopment {
		logger, loggerErr = zap.NewDevelopment(
			zap.WithClock(utcClock),
		)
	} else {
		logger, loggerErr = zap.NewProduction(
			zap.WithClock(utcClock),
		)
	}
	if loggerErr != nil {
		return nil, loggerErr
	}
	return &zapLogger{
		efficient: logger,
		sugared:   logger.Sugar(),
	}, nil
}
