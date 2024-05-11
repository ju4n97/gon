package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/go-chi/cors"
	dbsetup "github.com/mesatechlabs/gokit/internal/db"
	db "github.com/mesatechlabs/gokit/internal/db/codegen"

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

// Serve is a function that starts a web server with the provided configuration.
// It first checks the database connection and applies any pending database migrations.
//
// Example:
//
//	config := ServeConfig{
//	    Verbose: true,
//	    AllowedOrigins: []string{"http://example.com"},
//	    AllowedMethods: []string{"GET", "POST"},
//	    AllowedHeaders: []string{"Content-Type"},
//	    Port: "8080",
//	}
//
// Serve(config)
func Serve(config *ServeConfig) {
	err := dbsetup.NewDatabaseConnection(func(q *db.Queries) error {
		return nil
	})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	if err := dbsetup.NewDatabaseMigrations(); err != nil {
		log.Fatal("Failed to apply database migrations: ", err)
	}

	if config.Verbose {
		bytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(bytes))
	}

	r := InitRouter(&RouterConfig{
		Cors: cors.Options{
			AllowedOrigins:   config.AllowedOrigins,
			AllowedMethods:   config.AllowedMethods,
			AllowedHeaders:   config.AllowedHeaders,
			AllowCredentials: false,
			MaxAge:           300,
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
