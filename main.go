package main

import (
	"log"
	"os"

	"github.com/jm2097/gon/cmd"
	"github.com/jm2097/gon/internal/config"
	"github.com/urfave/cli/v2"
)

func main() {
	config.LoadConfigFromEnv()

	app := &cli.App{
		Name:                 "gon",
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
