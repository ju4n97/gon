package dbsetup

import (
	"context"
	"embed"
	"errors"
	"log"
	"path"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jm2097/gon/internal/codegen/db"
	"github.com/jm2097/gon/internal/config"
	"github.com/jm2097/gon/tools/logger"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embeddedMigrations embed.FS

type QueryFunc func(*db.Queries) error

func NewDatabaseConnection(queryFunc QueryFunc) error {
	logger.Log.Info().Msg("Checking database connection...")

	conn, err := pgx.Connect(context.Background(), config.Global.Postgres.Dsn("host", "port", "user", "password", "dbname", "sslmode"))
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

	defer func() {
		if err := conn.Close(context.Background()); err != nil {
			log.Default().Println(err)
		}
	}()

	queries := db.New(conn)

	return queryFunc(queries)
}

func NewDatabaseMigrations() error {
	logger.Log.Info().Msg("Checking database migrations...")

	goose.SetBaseFS(embeddedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	db, err := goose.OpenDBWithDriver("pgx", config.Global.Postgres.Dsn("host", "port", "user", "password", "dbname", "sslmode"))
	if err != nil {
		return err
	}

	defer db.Close()

	if err := goose.Up(db, path.Join("migrations")); err != nil {
		return err
	}

	return nil
}

// handleInvalidCatalogName (3D000) creates a new database once is verified that the database does not exist.
func handleInvalidCatalogName() error {
	logger.Log.Warn().Msg("Database does not exist. Creating it...")

	conn, err := pgx.Connect(context.Background(), config.Global.Postgres.Dsn("host", "port", "user", "password"))
	if err != nil {
		return err
	}

	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "CREATE DATABASE "+config.Global.Postgres.DBName)
	if err != nil {
		return err
	}

	return nil
}
