package dbsetup

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jm2097/gon/internal/codegen/db"
	"github.com/jm2097/gon/internal/config"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embeddedMigrations embed.FS

func NewDatabaseConnection(action func(*db.Queries) error) error {
	log.Default().Println("Checking database connection...")

	conn, err := pgx.Connect(context.Background(), config.AppConfig.PostgresDsn)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.InvalidCatalogName {
				return handleInvalidCatalogName()
			} else {
				return err
			}
		}

		return err
	}

	defer conn.Close(context.Background())

	queries := db.New(conn)

	return action(queries)
}

func NewDatabaseMigrations() error {
	log.Default().Println("Checking database migrations...")

	goose.SetBaseFS(embeddedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	db, err := goose.OpenDBWithDriver("pgx", config.AppConfig.PostgresDsn)
	if err != nil {
		return err
	}

	defer db.Close()

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	return nil
}

// handleInvalidCatalogName (3D000) creates a new database once is verified that the database does not exist.
func handleInvalidCatalogName() error {
	log.Default().Println("Database does not exist. Creating it...")

	conn, err := pgx.Connect(
		context.Background(),
		fmt.Sprintf("host=%s user=%s password=%s",
			config.AppConfig.PostgresHost,
			config.AppConfig.PostgresUser,
			config.AppConfig.PostgresPassword,
		),
	)
	if err != nil {
		return err
	}

	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "CREATE DATABASE "+config.AppConfig.PostgresDb)
	if err != nil {
		return err
	}

	return nil
}
