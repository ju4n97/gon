package cmd

import (
	"github.com/ju4n97/gon/apis"
	"github.com/urfave/cli/v2"
)

// NewServeCommand returns a new CLI command that starts a web server.
func NewServeCommand() *cli.Command {
	return &cli.Command{
		Name:  "serve",
		Usage: "Starts a web HTTP server to serve requests",
		Action: func(ctx *cli.Context) error {
			apis.Serve()
			return nil
		},
	}
}
