package apis

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/go-chi/cors"

	"net/http"
)

// ServeConfig defines the configuration for the web server
type ServeConfig struct {
	// HttpAddr is the address to listen on for HTTP connections (defaults to "127.0.0.1:3333")
	HttpAddr string

	// AllowedOrigins is an optional list of allowed CORS origins (defaults to "*")
	AllowedOrigins []string

	// AllowedMethods is an optional list of allowed HTTP methods (defaults to "GET,HEAD,POST,PUT,PATCH,DELETE")
	AllowedMethods []string

	// AllowedHeaders is an optional list of allowed HTTP headers (defaults to "User-Agent,Content-Type,Accept,Accept-Encoding,Accept-Language,Cache-Control,Connection,DNT,Host,Origin,Pragma,Referer")
	AllowedHeaders []string
}

// Serve starts a web server
func Serve(config ServeConfig) {
	// TODO: connect to the database

	// TODO: ensure the latest migrations are applied before starting the server
	// if err := runMigrations(); err != nil {
	// 	return nil, err
	// }

	fmt.Println(config)

	r := InitRouter(RouterConfig{
		Cors: cors.Options{
			AllowedOrigins:   config.AllowedOrigins,
			AllowedMethods:   config.AllowedMethods,
			AllowedHeaders:   config.AllowedHeaders,
			AllowCredentials: true,
		},
		Timeout: 60 * time.Second,
	})

	protocol := "http"
	version := "v1"

	c := color.New()
	c.Printf("└─ REST API: %s\n", color.GreenString("%s://%s/api/%s",
		protocol,
		config.HttpAddr,
		version,
	))

	if err := http.ListenAndServe(config.HttpAddr, r); err != nil {
		log.Fatal(err)
	}
}
