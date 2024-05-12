package cmd

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/jm2097/gon/apis"
	"github.com/urfave/cli/v2"
)

// NewServeCommand returns a new CLI command that starts a web server.
func NewServeCommand() *cli.Command {
	defaultPort := 3000

	defaultOrigins := []string{"*"}

	defaultMethods := []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
	}

	defaultHeaders := []string{
		"User-Agent",
		"Content-Type",
		"Accept",
		"Accept-Encoding",
		"Accept-Language",
		"Cache-Control",
		"Connection",
		"DNT",
		"Host",
		"Origin",
		"Pragma",
		"Referer",
	}

	return &cli.Command{
		Name:  "serve",
		Usage: "Starts a web HTTP server to serve requests",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Usage:       "Port where the server will listen for HTTP requests",
				Value:       defaultPort,
				DefaultText: strconv.Itoa(defaultPort),
				EnvVars:     []string{"PORT"},
			},
			&cli.StringSliceFlag{
				Name:        "allowed-origins",
				Usage:       "List of domains allowed to access the server via CORS",
				Value:       cli.NewStringSlice(defaultOrigins...),
				DefaultText: strings.Join(defaultOrigins, ","),
				EnvVars:     []string{"CORS_ALLOWED_ORIGINS"},
			},
			&cli.StringSliceFlag{
				Name:        "allowed-methods",
				Usage:       "List of HTTP methods allowed to access the server",
				Value:       cli.NewStringSlice(defaultMethods...),
				DefaultText: strings.Join(defaultMethods, ","),
				EnvVars:     []string{"CORS_ALLOWED_METHODS"},
			},
			&cli.StringSliceFlag{
				Name:        "allowed-headers",
				Usage:       "List of HTTP headers allowed to access the server",
				Value:       cli.NewStringSlice(defaultHeaders...),
				DefaultText: strings.Join(defaultHeaders, ","),
				EnvVars:     []string{"CORS_ALLOWED_HEADERS"},
			},
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"v"},
				Usage:       "Prints the configuration used to start the server",
				DefaultText: "false",
			},
		},
		Action: func(ctx *cli.Context) error {
			apis.Serve(&apis.ServeConfig{
				Port:           ctx.String("port"),
				AllowedOrigins: ctx.StringSlice("allowed-origins"),
				AllowedMethods: ctx.StringSlice("allowed-methods"),
				AllowedHeaders: ctx.StringSlice("allowed-headers"),
				Verbose:        ctx.Bool("verbose"),
			})

			return nil
		},
	}
}
