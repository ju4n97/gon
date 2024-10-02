package config

import (
	"fmt"

	"github.com/ju4n97/gon/internal/env"
	"github.com/ju4n97/gon/tools/defaults"
	"github.com/joho/godotenv"
)

type EnvConfigLoader struct{}

// LoadConfig loads the global configuration from environment variables using a .env file.
func (l *EnvConfigLoader) LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Error loading .env file: %v", err)
	}

	Global = &Config{
		App: AppConfig{
			Name: env.New("APP_NAME").MustToString(),
			Env:  Environment(env.New("APP_ENV").MustToString()),
		},
		Server: ServerConfig{
			Port:           env.New("PORT").WithDefault(defaults.ServerPort).MustToInt(),
			Host:           env.New("HOST").WithDefault(defaults.ServerHost).MustToString(),
			AllowedOrigins: env.New("CORS_ALLOWED_ORIGINS").WithDefault(defaults.ServerAllowedOrigins).MustToStringSlice(","),
			AllowedMethods: env.New("CORS_ALLOWED_METHODS").WithDefault(defaults.ServerAllowedMethods).MustToStringSlice(","),
			AllowedHeaders: env.New("CORS_ALLOWED_HEADERS").WithDefault(defaults.ServerAllowedHeaders).MustToStringSlice(","),
		},
		Postgres: PostgresConfig{
			Host:     env.New("POSTGRES_HOST").MustToString(),
			Port:     env.New("POSTGRES_PORT").MustToInt(),
			User:     env.New("POSTGRES_USER").MustToString(),
			Password: env.New("POSTGRES_PASSWORD").MustToString(),
			DBName:   env.New("POSTGRES_DBNAME").MustToString(),
			SslMode:  env.New("POSTGRES_SSL_MODE").MustToString(),
		},
		Logger: LoggerConfig{
			IsWriteToStdoutEnabled:   env.New("LOGGER_TO_STDOUT_ENABLED").WithDefault(defaults.LoggerWriteToStdoutEnabled).MustToBool(),
			IsWriteToFileEnabled:     env.New("LOGGER_TO_FILE_ENABLED").WithDefault(defaults.LoggerWriteToStdoutEnabled).MustToBool(),
			MinLevel:                 env.New("LOGGER_MIN_LEVEL").WithDefault(defaults.LoggerMinLevel).MustToString(),
			FilePath:                 env.New("LOGGER_FILE_PATH").WithDefault(defaults.LoggerFilePath).MustToString(),
			FileMaxSize:              env.New("LOGGER_FILE_MAX_SIZE_IN_MB").WithDefault(defaults.LoggerMaxSize).MustToInt(),
			FileMaxAge:               env.New("LOGGER_FILE_MAX_AGE").WithDefault(defaults.LoggerMaxAge).MustToInt(),
			FileMaxBackups:           env.New("LOGGER_FILE_MAX_BACKUPS").WithDefault(defaults.LoggerMaxBackups).MustToInt(),
			IsFileLocalTimeEnabled:   env.New("LOGGER_FILE_LOCAL_TIME_ENABLED").WithDefault(defaults.LoggerLocalTimeEnabled).MustToBool(),
			IsFileCompressionEnabled: env.New("LOGGER_FILE_COMPRESS_ENABLED").WithDefault(defaults.LoggerCompressionEnabled).MustToBool(),
			PrettyPrintEnvironments:  env.New("LOGGER_PRETTY_PRINT_ENVIRONMENTS").WithDefault(defaults.LoggerPrettyPrintEnvironments).MustToStringSlice(","),
		},
	}

	if err := Global.Validate(); err != nil {
		return nil, fmt.Errorf("Invalid config: %v", err)
	}

	return Global, nil
}
