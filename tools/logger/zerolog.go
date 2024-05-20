package logger

import (
	"io"
	"os"
	"slices"

	"github.com/jm2097/gon/internal/config"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZeroLogger struct {
	logger zerolog.Logger
}

type ZeroLogEvent struct {
	event *zerolog.Event
}

// NewZeroLogger creates a new instance of the ZeroLogger struct.
//
// It initializes the ZeroLogger with the appropriate log level and log writers based on the global configuration.
// The log level is set to the minimum level specified in the configuration, or defaults to info if it cannot be parsed.
// The log writers include standard output and/or a file writer based on the configuration.
func NewZeroLogger() Logger {
	var writers []io.Writer

	if config.Global.App.IsProd() {
		logLevel, err := zerolog.ParseLevel(config.Global.Logger.MinLevel)
		if err != nil {
			logLevel = zerolog.InfoLevel
		}

		zerolog.SetGlobalLevel(logLevel)
	}

	if config.Global.Logger.IsWriteToStdoutEnabled {
		if slices.Contains(config.Global.Logger.PrettyPrintEnvironments, string(config.Global.App.Env)) {
			writers = append(writers, zerolog.ConsoleWriter{Out: os.Stdout})
		} else {
			writers = append(writers, os.Stdout)
		}
	}

	if config.Global.Logger.IsWriteToFileEnabled {
		writers = append(writers, &lumberjack.Logger{
			Filename:   config.Global.Logger.FilePath,
			MaxSize:    config.Global.Logger.FileMaxSize,
			MaxAge:     config.Global.Logger.FileMaxAge,
			MaxBackups: config.Global.Logger.FileMaxBackups,
			LocalTime:  config.Global.Logger.IsFileLocalTimeEnabled,
			Compress:   config.Global.Logger.IsFileCompressionEnabled,
		})
	}

	multi := io.MultiWriter(writers...)

	zl := zerolog.New(multi).With().Timestamp().Caller().Logger()

	Log = &ZeroLogger{logger: zl}

	return Log
}

func (zl *ZeroLogger) Debug() LogEvent {
	return &ZeroLogEvent{event: zl.logger.Debug()}
}

func (zl *ZeroLogger) Info() LogEvent {
	return &ZeroLogEvent{event: zl.logger.Info()}
}

func (zl *ZeroLogger) Warn() LogEvent {
	return &ZeroLogEvent{event: zl.logger.Warn()}
}

func (zl *ZeroLogger) Error() LogEvent {
	return &ZeroLogEvent{event: zl.logger.Error()}
}

func (zl *ZeroLogger) Fatal() LogEvent {
	return &ZeroLogEvent{event: zl.logger.Fatal()}
}

func (zl *ZeroLogger) Panic() LogEvent {
	return &ZeroLogEvent{event: zl.logger.Panic()}
}

func (zle *ZeroLogEvent) WithFields(fields Fields) LogEvent {
	zle.event = zle.event.Fields(map[string]interface{}(fields))
	return zle
}

func (zle *ZeroLogEvent) Msg(msg string) {
	zle.event.Msg(msg)
}
