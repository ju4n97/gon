package defaults

import "net/http"

const (
	AppName = "gon"
)

const (
	ServerHost           = "127.0.0.1"
	ServerPort           = "3000"
	ServerAllowedOrigins = "*"
)

var (
	ServerAllowedMethods = [...]string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
	}
	ServerAllowedHeaders = [...]string{
		"User-Agent",
		"Content-Type",
		"Accept",
		"Accept-Encoding",
		"Accept-Language",
		"Cache-Control",
		"Connection",
		"DNT",
		"Host",
		"Origin",
		"Pragma",
		"Referer",
	}
)

const (
	PostgresSslMode = "disable"
)

const (
	LoggerWriteToStdoutEnabled = true
	LoggerWriteToFileEnabled   = false
	LoggerMinLevel             = "debug"
	LoggerFilePath             = "logs/app.log"
	LoggerMaxSize              = 10 // in megabytes
	LoggerMaxAge               = 7  // in days
	LoggerMaxBackups           = 3
	LoggerLocalTimeEnabled     = false
	LoggerCompressionEnabled   = false
)

var LoggerPrettyPrintEnvironments = []string{"dev", "test"}
