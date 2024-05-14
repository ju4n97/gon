package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	zapLogger *zap.Logger
}

var Log *Logger

// Debug (-1) records messages useful for debugging.
func (l *Logger) Debug(message string, fields ...zap.Field) {
	l.zapLogger.Debug(message, fields...)
}

// Info (0) records standard application events and usage statistics.
func (l *Logger) Info(message string, fields ...zap.Field) {
	l.zapLogger.Info(message, fields...)
}

// Warn (1) records unusual events that need attention before they escalate to more sever issues.
func (l *Logger) Warn(message string, fields ...zap.Field) {
	l.zapLogger.Warn(message, fields...)
}

// Error (2) records unexpected error conditions in the application.
func (l *Logger) Error(err error, fields ...zap.Field) {
	l.zapLogger.Error(err.Error(), fields...)
}

// DPanic (3) records sever error conditions in development. It behaves like panic (4) in development and error (2) in production.
func (l *Logger) DPanic(err error, fields ...zap.Field) {
	l.zapLogger.DPanic(err.Error(), fields...)
}

// Panic (4) calls panic() after logging an error condition.
func (l *Logger) Panic(err error, fields ...zap.Field) {
	l.zapLogger.Panic(err.Error(), fields...)
}

// Fatal (5) calls os.Exit(1) after logging an error condition.
func (l *Logger) Fatal(err error, fields ...zap.Field) {
	l.zapLogger.Fatal(err.Error(), fields...)
}

func NewLogger() *zap.Logger {
	// stdout := zapcore.AddSync(os.Stdout)

	// file := zapcore.AddSync(&lumberjack.Logger{
	// 	Filename:   config.Config.Logs.File,
	// 	MaxSize:    config.Config.Logs.MaxSize,
	// 	MaxAge:     config.Config.Logs.MaxAge,
	// 	MaxBackups: config.Config.Logs.MaxBackups,
	// })

	// return zap.New(zapLogger)

	return zap.NewExample()
}
