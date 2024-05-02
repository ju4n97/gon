package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/go-chi/cors"

	"net/http"
)

// ServeConfig defines the configuration for the web server
type ServeConfig struct {
	// Port is the HTTP port to listen on (default: 3000)
	Port string

	// AllowedOrigins is an optional list of allowed CORS origins (default: "*")
	AllowedOrigins []string

	// AllowedMethods is an optional list of allowed HTTP methods (default: "GET,HEAD,POST,PUT,PATCH,DELETE")
	AllowedMethods []string

	// AllowedHeaders is an optional list of allowed HTTP headers (default: "User-Agent,Content-Type,Accept,Accept-Encoding,Accept-Language,Cache-Control,Connection,DNT,Host,Origin,Pragma,Referer")
	AllowedHeaders []string

	// Verbose enables verbose logging during the server startup (default: false)
	Verbose bool
}

// Serve starts a web server
func Serve(config ServeConfig) {
	// TODO: connect to the database

	// TODO: ensure the latest migrations are applied before starting the server
	// if err := runMigrations(); err != nil {
	// 	return nil, err
	// }

	if config.Verbose {
		bytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(bytes))
	}

	r := InitRouter(RouterConfig{
		Cors: cors.Options{
			AllowedOrigins:   config.AllowedOrigins,
			AllowedMethods:   config.AllowedMethods,
			AllowedHeaders:   config.AllowedHeaders,
			AllowCredentials: true,
		},
		Timeout: 60 * time.Second,
	})

	httpAddr := fmt.Sprintf("127.0.0.1:%s", config.Port)

	c := color.New()
	c.Printf("└─ REST API: %s\n", color.GreenString("http://%s/api", httpAddr))
	c.Printf("   ├─ Health: %s\n", color.GreenString("http://%s/api/health", httpAddr))
	c.Printf("   └─ v1: %s\n", color.GreenString("http://%s/api/v1", httpAddr))

	if err := http.ListenAndServe(httpAddr, r); err != nil {
		log.Fatal(err)
	}
}
