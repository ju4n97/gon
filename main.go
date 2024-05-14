package main

import (
	"log"
	"os"

	"github.com/jm2097/gon/cmd"
	"github.com/jm2097/gon/internal/config"
	"github.com/urfave/cli/v2"
)

func main() {
	factory, err := config.NewConfigLoaderFactory("env")
	if err != nil {
		log.Fatal(err)
	}

	config, err := factory.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name:                 config.App.Name,
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
