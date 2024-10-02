package apis

import (
	"net/http"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/ju4n97/gon/internal/codegen/db"
	"github.com/ju4n97/gon/internal/config"
	dbsetup "github.com/ju4n97/gon/internal/db"
	"github.com/ju4n97/gon/tools/logger"
	"github.com/labstack/echo/v4"
)

func Serve() {
	checkDatabase()

	router := NewRouter()

	httpAddr := config.Global.Server.Host + ":" + strconv.Itoa(config.Global.Server.Port)

	printRoutes("http://"+httpAddr, router.Routes())

	server := &http.Server{
		Addr:              httpAddr,
		Handler:           router,
		ReadHeaderTimeout: 3 * time.Second,
	}

	if err := router.StartServer(server); err != nil {
		logger.Log.Fatal().WithFields(logger.Fields{"error": err}).Msg("Failed to start the server")
	}
}

func checkDatabase() {
	logger.Log.Info().Msg("Checking database connection...")

	err := dbsetup.NewDatabaseConnection(func(q *db.Queries) error {
		return nil
	})
	if err != nil {
		logger.Log.Fatal().WithFields(logger.Fields{"error": err}).Msg("Failed to connect to the database")
	}

	logger.Log.Info().Msg("Checking database migrations...")

	if err := dbsetup.NewDatabaseMigrations(); err != nil {
		logger.Log.Fatal().WithFields(logger.Fields{"error": err}).Msg("Failed to run database migrations")
	}
}

func printRoutes(httpAddr string, routes []*echo.Route) {
	c := color.New()
	c.Printf("└─ REST API\n")

	for idx, route := range routes {
		threeSymbol := "  ├─"
		if idx == len(routes)-1 {
			threeSymbol = "  └─"
		}

		c.Printf("%s %s: %s\n", threeSymbol, color.MagentaString(route.Method), color.GreenString(httpAddr+route.Path))
	}
}
