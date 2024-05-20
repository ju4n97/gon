package config

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jm2097/gon/tools/custom_validator"
)

type Environment string

const (
	Dev  Environment = "dev"
	Test Environment = "test"
	Prod Environment = "prod"
)

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Postgres PostgresConfig
	Logger   LoggerConfig
}

type ConfigLoader interface {
	LoadConfig() (*Config, error)
}

var Global *Config

type AppConfig struct {
	Name string      `validate:"required"`
	Env  Environment `validate:"required,oneof=dev test prod"`
}

type ServerConfig struct {
	Port           int      `validate:"omitempty,port"`
	Host           string   `validate:"omitempty"`
	AllowedOrigins []string `validate:"dive,omitempty,eq=*|url"`
	AllowedMethods []string `validate:"dive,omitempty,eq=*|oneof=GET HEAD POST PUT PATCH DELETE OPTIONS CONNECT TRACE"`
	AllowedHeaders []string `validate:"omitempty"`
}

type PostgresConfig struct {
	Host     string `validate:"required"`
	Port     int    `validate:"required,port"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	DBName   string `validate:"required"`
	SslMode  string `validate:"omitempty,oneof=disable allow prefer require verify-ca verify-full"`
}

type LoggerConfig struct {
	IsWriteToStdoutEnabled   bool     `validate:"omitempty"`
	IsWriteToFileEnabled     bool     `validate:"omitempty"`
	MinLevel                 string   `validate:"omitempty,oneof=debug info warn error fatal panic"`
	FilePath                 string   `validate:"omitempty"`
	FileMaxSize              int      `validate:"omitempty,gt=0"` // in megabytes
	FileMaxAge               int      `validate:"omitempty,gt=0"` // in days
	FileMaxBackups           int      `validate:"omitempty,gt=0"`
	IsFileLocalTimeEnabled   bool     `validate:"omitempty"`
	IsFileCompressionEnabled bool     `validate:"omitempty"`
	PrettyPrintEnvironments  []string `validate:"dive,omitempty,oneof=dev test prod"`
}

func (c *Config) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.RegisterValidation("port", custom_validator.ValidatePort); err != nil {
		return err
	}

	return validate.Struct(c)
}

func (c *AppConfig) IsDev() bool {
	return c.Env == Dev
}

func (c *AppConfig) IsTest() bool {
	return c.Env == Test
}

func (c *AppConfig) IsProd() bool {
	return c.Env == Prod
}

func (c *PostgresConfig) Dsn(includedProps ...string) string {
	props := map[string]string{
		"host":     c.Host,
		"port":     strconv.Itoa(c.Port),
		"user":     c.User,
		"password": c.Password,
		"dbname":   c.DBName,
		"sslmode":  c.SslMode,
	}

	var dsn string
	for _, prop := range includedProps {
		dsn += fmt.Sprintf("%s=%s ", prop, props[prop])
	}

	return strings.TrimSpace(dsn)
}
