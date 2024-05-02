package main

import (
	"log"
	"os"

	"github.com/mesatechlabs/kitten/cmd"
	"github.com/mesatechlabs/kitten/internals/config"
	"github.com/urfave/cli/v2"
)

func main() {
	config.LoadEnv()

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
