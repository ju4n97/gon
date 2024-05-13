package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jm2097/gon/tools/validators"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string `validate:"omitempty,port"`

	CorsAllowedOrigins []string `validate:"dive,omitempty,eq=*|url"`
	CorsAllowedMethods []string `validate:"dive,omitempty,eq=*|oneof=GET HEAD POST PUT PATCH DELETE OPTIONS CONNECT TRACE"`
	CorsAllowedHeaders []string

	PostgresHost     string `validate:"required"`
	PostgresPort     string `validate:"required,port"`
	PostgresUser     string `validate:"required"`
	PostgresPassword string `validate:"required"`
	PostgresDb       string `validate:"required"`
	PostgresSsl      bool   `validate:"omitempty,boolean"`
	PostgresDsn      string // The Data Source Name (Dsn) is not passed in the environment. It is calculated from the other fields.
}

var AppConfig *Config

// LoadConfigFromEnv loads the environment variables from the .env file, validates them for correctness and sets
// the value for the global AppConfig variable so it can be used in other parts of the application
// with the values already converted to the correct type.
func LoadConfigFromEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	AppConfig = &Config{
		Port: os.Getenv("PORT"),

		CorsAllowedOrigins: strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
		CorsAllowedMethods: strings.Split(os.Getenv("CORS_ALLOWED_METHODS"), ","),
		CorsAllowedHeaders: strings.Split(os.Getenv("CORS_ALLOWED_HEADERS"), ","),

		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDb:       os.Getenv("POSTGRES_DB"),
		PostgresSsl:      os.Getenv("POSTGRES_SSL") == "true",
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.RegisterValidation("port", validators.ValidatePort); err != nil {
		log.Fatal(err)
	}

	if err := validate.Struct(AppConfig); err != nil {
		log.Fatal(err)
	}

	setPostgresDsn(AppConfig)
}

// setPostgresDsn calculates the Data Source Name (Dsn) from the other database fields
// of the AppConfig struct and mutates the AppConfig.PostgresDsn field directly.
func setPostgresDsn(config *Config) {
	dsnQuery := "?"

	if !config.PostgresSsl {
		dsnQuery += "sslmode=disable"
	}

	AppConfig.PostgresDsn = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s%s",
		AppConfig.PostgresUser,
		AppConfig.PostgresPassword,
		AppConfig.PostgresHost,
		AppConfig.PostgresPort,
		AppConfig.PostgresDb,
		dsnQuery,
	)
}
