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
	LogsEnabled     = true
	LogsFile        = "logs/app.log"
	LogsMaxSizeInMb = 10 // megabytes
	LogsMaxAge      = 7  // days
	LogsMaxBackups  = 3
)
