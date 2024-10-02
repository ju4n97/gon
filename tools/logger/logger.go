package logger

type Fields map[string]interface{}

type LogEvent interface {
	WithFields(fields Fields) LogEvent
	Msg(msg string)
}

type Logger interface {
	Debug() LogEvent
	Info() LogEvent
	Warn() LogEvent
	Error() LogEvent
	Fatal() LogEvent
	Panic() LogEvent
}

var Log Logger
