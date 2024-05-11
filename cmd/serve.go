package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/mesatechlabs/gokit/apis"
	"github.com/urfave/cli/v2"
)

// NewServeCommand returns a new CLI command that starts a web server
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
		Usage: fmt.Sprintf("Starts a web HTTP server on http://127.0.0.1:%s", strconv.Itoa(defaultPort)),
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
		Action: func(c *cli.Context) error {
			apis.Serve(&apis.ServeConfig{
				Port:           c.String("port"),
				AllowedOrigins: c.StringSlice("allowed-origins"),
				AllowedMethods: c.StringSlice("allowed-methods"),
				AllowedHeaders: c.StringSlice("allowed-headers"),
				Verbose:        c.Bool("verbose"),
			})
			return nil
		},
	}
}
