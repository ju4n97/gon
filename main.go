package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mesatechlabs/kitten/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	app := &cli.App{
		Name:                 "kitten",
		Usage:                "SvelteKit template powered by a Go backend",
		Version:              "0.0.1",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			cmd.NewServeCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
