package config

import (
	"log"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/mesatechlabs/kitten/validators"
)

type EnvConfig struct {
	Port string `validate:"omitempty,port"`

	CorsAllowedOrigins []string `validate:"dive,omitempty,eq=*|url"`
	CorsAllowedMethods []string `validate:"dive,omitempty,eq=*|oneof=GET HEAD POST PUT PATCH DELETE OPTIONS CONNECT TRACE"`
	CorsAllowedHeaders []string

	PostgresHost     string `validate:"required"`
	PostgresPort     string `validate:"required,port"`
	PostgresUser     string `validate:"required"`
	PostgresPassword string `validate:"required"`
	PostgresDBName   string `validate:"required"`
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	config := &EnvConfig{
		Port: os.Getenv("PORT"),

		CorsAllowedOrigins: strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
		CorsAllowedMethods: strings.Split(os.Getenv("CORS_ALLOWED_METHODS"), ","),
		CorsAllowedHeaders: strings.Split(os.Getenv("CORS_ALLOWED_HEADERS"), ","),

		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDBName:   os.Getenv("POSTGRES_DB_NAME"),
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("port", validators.ValidatePort)
	if err := validate.Struct(config); err != nil {
		log.Fatal(err)
	}
}
