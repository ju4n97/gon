package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mesatechlabs/kitten/apis"
	"github.com/urfave/cli/v2"
)

// NewServeCommand returns a new CLI command that starts a web server
func NewServeCommand() *cli.Command {
	defaultHttpAddr := "127.0.0.1:3000"

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
		Usage: fmt.Sprintf("Starts a web HTTP server (default: %s)", defaultHttpAddr),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "http-addr",
				Usage:       "Address and port where the server will listen for HTTP requests",
				Value:       defaultHttpAddr,
				DefaultText: defaultHttpAddr,
				EnvVars:     []string{"HTTP_ADDR"},
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
		},
		Action: func(c *cli.Context) error {
			apis.Serve(apis.ServeConfig{
				HttpAddr:       c.String("http-addr"),
				AllowedOrigins: c.StringSlice("allowed-origins"),
				AllowedMethods: c.StringSlice("allowed-methods"),
				AllowedHeaders: c.StringSlice("allowed-headers"),
			})
			return nil
		},
	}
}
