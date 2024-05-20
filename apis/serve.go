package apis

import (
	"net/http"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/go-chi/cors"
	db "github.com/jm2097/gon/internal/codegen/db"
	"github.com/jm2097/gon/internal/config"
	dbsetup "github.com/jm2097/gon/internal/db"
	"github.com/jm2097/gon/tools/logger"
)

func Serve() {
	err := dbsetup.NewDatabaseConnection(func(q *db.Queries) error {
		return nil
	})
	if err != nil {
		logger.Log.Fatal().WithFields(logger.Fields{"error": err}).Msg("Failed to connect to the database")
	}

	if err := dbsetup.NewDatabaseMigrations(); err != nil {
		logger.Log.Fatal().WithFields(logger.Fields{"error": err}).Msg("Failed to run database migrations")
	}

	router := InitRouter(&RouterConfig{
		Cors: cors.Options{
			AllowedOrigins:   config.Global.Server.AllowedOrigins,
			AllowedMethods:   config.Global.Server.AllowedMethods,
			AllowedHeaders:   config.Global.Server.AllowedHeaders,
			AllowCredentials: false,
			MaxAge:           300,
		},
		Timeout: 60 * time.Second,
		RateLimit: RouterConfigRateLimit{
			RequestsLimit:  100,
			RequestsWindow: 5 * time.Minute,
		},
	})

	httpAddr := config.Global.Server.Host + ":" + strconv.Itoa(config.Global.Server.Port)
	fullHttpAddr := "http://" + httpAddr

	c := color.New()
	c.Printf("└─ REST API: %s\n", color.GreenString(fullHttpAddr+"/api"))
	c.Printf("   ├─ Health: %s\n", color.GreenString(fullHttpAddr+"/health"))
	c.Printf("   └─ v1: %s\n", color.GreenString(fullHttpAddr+"/api/v1"))

	server := &http.Server{
		Addr:              httpAddr,
		Handler:           router,
		ReadHeaderTimeout: 3 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Log.Fatal().WithFields(logger.Fields{"error": err}).Msg("Failed to start the server")
	}
}
