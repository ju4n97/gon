package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jm2097/gon/internal/config"
)

func main() {
	config.LoadConfigFromEnv()

	conn, err := pgx.Connect(context.Background(), config.AppConfig.PostgresDsn)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	if _, err = conn.Exec(context.Background(), "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"); err != nil {
		log.Panic(err)
	}

	log.Println("Database reset successfully")
}
