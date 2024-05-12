package dbsetup

import (
	"context"
	"log"
	"os"

	"ariga.io/atlas-go-sdk/atlasexec"
	"github.com/jackc/pgx/v5"
	"github.com/jm2097/gon/internal/config"
	db "github.com/jm2097/gon/internal/db/codegen"
)

type DBAction func(*db.Queries) error

func NewDatabaseConnection(action DBAction) error {
	log.Default().Println("Checking database connection...")

	conn, err := pgx.Connect(context.Background(), config.AppConfig.PostgresDsn)
	if err != nil {
		return err
	}

	defer conn.Close(context.Background())

	queries := db.New(conn)

	return action(queries)
}

func NewDatabaseMigrations() error {
	log.Default().Println("Checking database migrations...")

	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithMigrations(
			os.DirFS("./internal/db/migrations"),
		),
	)
	if err != nil {
		return err
	}

	defer workdir.Close()

	client, err := atlasexec.NewClient(workdir.Path(), "atlas")
	if err != nil {
		return err
	}

	status, err := client.MigrateStatus(context.Background(), &atlasexec.MigrateStatusParams{
		URL: os.Getenv("DATABASE_URI"),
	})
	if err != nil {
		return err
	}

	if status.Status == "OK" {
		log.Default().Printf("Database is up-to-date and no migrations need to be applied  (current version: %s\n)", status.Current)

		return nil
	}

	res, err := client.MigrateApply(context.Background(), &atlasexec.MigrateApplyParams{
		URL: os.Getenv("DATABASE_URI"),
	})
	if err != nil {
		return err
	}

	log.Default().Printf("Applied %d migrations\n", len(res.Applied))

	return nil
}
