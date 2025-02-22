package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	// DebugF logs a formatted debug message.
	//
	// Debug messages are typically used for debugging and diagnostic purposes in development environments.
	DebugF(format string, args ...interface{})

	// InfoF logs a formatted informational message.
	InfoF(format string, args ...interface{})

	// WarnF logs a formatted warning message.
	WarnF(format string, args ...interface{})

	// ErrorF logs a formatted error message.
	ErrorF(format string, args ...interface{})

	// FatalF logs a formatted fatal message and then exits the application.
	FatalF(format string, args ...interface{})

	// DebugWithData logs a debug message with additional data.
	//
	// Debug messages are typically used for debugging and diagnostic purposes in development environments.
	DebugWithData(message string, data AdditionalData)

	// InfoWithData logs an informational message with additional data.
	InfoWithData(message string, data AdditionalData)

	// WarnWithData logs a warning message with additional data.
	WarnWithData(message string, data AdditionalData)

	// ErrorWithData logs an error message with additional data.
	ErrorWithData(message string, data AdditionalData)

	// FatalWithData logs a fatal message with additional data and then exits the application.
	FatalWithData(message string, data AdditionalData)

	// Close flushes any buffered log entries.
	Close() error
}

type AdditionalData map[string]interface{}

type zapLogger struct {
	efficient *zap.Logger
	sugared   *zap.SugaredLogger
}

func (l *zapLogger) DebugF(format string, args ...interface{}) {
	l.log(zapcore.DebugLevel, format, args...)
}

func (l *zapLogger) InfoF(format string, args ...interface{}) {
	l.log(zapcore.InfoLevel, format, args...)
}

func (l *zapLogger) WarnF(format string, args ...interface{}) {
	l.log(zapcore.WarnLevel, format, args...)
}

func (l *zapLogger) ErrorF(format string, args ...interface{}) {
	l.log(zapcore.ErrorLevel, format, args...)
}

func (l *zapLogger) FatalF(format string, args ...interface{}) {
	l.log(zapcore.FatalLevel, format, args...)
}

func (l *zapLogger) log(level zapcore.Level, format string, args ...interface{}) {
	l.sugared.Logf(level, format, args...)
}

func (l *zapLogger) DebugWithData(message string, data AdditionalData) {
	l.logWithData(zapcore.DebugLevel, message, data)
}

func (l *zapLogger) InfoWithData(message string, data AdditionalData) {
	l.logWithData(zapcore.InfoLevel, message, data)
}

func (l *zapLogger) WarnWithData(message string, data AdditionalData) {
	l.logWithData(zapcore.WarnLevel, message, data)
}

func (l *zapLogger) ErrorWithData(message string, data AdditionalData) {
	l.logWithData(zapcore.ErrorLevel, message, data)
}

func (l *zapLogger) FatalWithData(message string, data AdditionalData) {
	l.logWithData(zapcore.FatalLevel, message, data)
}

func (l *zapLogger) logWithData(level zapcore.Level, message string, data AdditionalData) {
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
