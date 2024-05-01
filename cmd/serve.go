package cmd

import (
	"github.com/mesatechlabs/kitten/apis"
	"github.com/spf13/cobra"
)

// NewServeCommand creates a new cobra.Command for starting the web server.
func NewServeCommand() *cobra.Command {
	var httpAddr string
	var allowedOrigins []string
	var allowedMethods []string
	var allowedHeaders []string

	var command = &cobra.Command{
		Use:   "serve",
		Short: "Starts the web server (default to 127.0.0.1:3333)",
		Run: func(cmd *cobra.Command, args []string) {
			apis.Serve(apis.ServeConfig{
				HttpAddr:       httpAddr,
				AllowedOrigins: allowedOrigins,
				AllowedMethods: allowedMethods,
				AllowedHeaders: allowedHeaders,
			})
		},
	}

	command.PersistentFlags().StringVar(
		&httpAddr,
		"http-addr",
		"127.0.0.1:3333",
		"Address and port where the server will listen for HTTP requests",
	)

	command.PersistentFlags().StringSliceVar(
		&allowedOrigins,
		"allowed-origins",
		allowedOrigins,
		"List of domains allowed to access the server via CORS",
	)

	command.PersistentFlags().StringSliceVar(
		&allowedMethods,
		"allowed-methods",
		allowedMethods,
		"List of HTTP methods allowed to access the server",
	)

	command.PersistentFlags().StringSliceVar(
		&allowedHeaders,
		"allowed-headers",
		allowedHeaders,
		"List of HTTP headers allowed to access the server",
	)

	return command
}
