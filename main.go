package main

import (
	"log"
	"os"

	"github.com/ju4n97/gon/cmd"
	"github.com/ju4n97/gon/internal/config"
	"github.com/ju4n97/gon/tools/logger"
	"github.com/urfave/cli/v2"
)

func main() {
	factory, err := config.NewConfigLoaderFactory("env")
	if err != nil {
		log.Fatalf("Failed to create config loader factory: %s", err)
	}

	config, err := factory.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	logger.Log = logger.NewZeroLogger()
	logger.Log.Info().WithFields(logger.Fields{"env": config.App.Env}).Msg("Starting " + config.App.Name)

	app := &cli.App{
		Name:                 config.App.Name,
		Usage:                "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		Version:              "0.0.1",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			cmd.NewServeCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.Log.Error().WithFields(logger.Fields{"error": err}).Msg("Failed to start the application")
	}
}
