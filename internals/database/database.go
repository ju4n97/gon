package database

import (
	"context"
	"fmt"
	"os"

	"ariga.io/atlas-go-sdk/atlasexec"
	"github.com/jackc/pgx/v5"
)

func CheckDatabaseConnection() error {
	fmt.Println("Checking database connection...")

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URI"))
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	return nil
}

func CheckDatabaseMigrations() error {
	fmt.Println("Checking database migrations...")

	// Define the execution context, supplying a migration directory
	// and potentially an `atlas.hcl` configuration file using `atlasexec.WithHCL`
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithMigrations(
			os.DirFS("./internals/database/migrations"),
		),
	)
	if err != nil {
		return err
	}
	defer workdir.Close()

	// Initialize the atlas client
	client, err := atlasexec.NewClient(workdir.Path(), "atlas")
	if err != nil {
		return err
	}

	status, err := client.MigrateStatus(context.Background(), &atlasexec.MigrateStatusParams{
		URL: os.Getenv("DATABASE_URI"),
	})
	fmt.Println(status)
	if err != nil {
		return err
	}
	if status.Status == "OK" {
		fmt.Printf("Database is up-to-date and no migrations need to be applied  (current version: %s\n)", status.Current)
		return nil
	}

	// Run `atlas migrate apply` on the PostgreSQL database
	res, err := client.MigrateApply(context.Background(), &atlasexec.MigrateApplyParams{
		URL: os.Getenv("DATABASE_URI"),
	})
	if err != nil {
		return err
	}
	fmt.Printf("Applied %d migrations\n", len(res.Applied))

	return nil
}
