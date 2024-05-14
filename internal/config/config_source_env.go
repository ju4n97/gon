package config

import (
	"fmt"

	"github.com/jm2097/gon/internal/env"
	"github.com/jm2097/gon/tools/defaults"
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
		Logs: LogsConfig{
			IsEnabled:   env.New("LOGS_ENABLED").WithDefault(defaults.LogsEnabled).MustToBool(),
			File:        env.New("LOGS_FILE").WithDefault(defaults.LogsFile).MustToString(),
			MaxSizeInMb: env.New("LOGS_MAX_SIZE_IN_MB").WithDefault(defaults.LogsMaxSizeInMb).MustToFloat32(),
			MaxAge:      env.New("LOGS_MAX_AGE").WithDefault(defaults.LogsMaxAge).MustToInt(),
			MaxBackups:  env.New("LOGS_MAX_BACKUPS").WithDefault(defaults.LogsMaxBackups).MustToInt(),
		},
	}

	if err := Global.Validate(); err != nil {
		return nil, fmt.Errorf("Invalid config: %v", err)
	}

	return Global, nil
}
