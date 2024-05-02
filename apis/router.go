package apis

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"

	"github.com/mesatechlabs/kitten/apis/health"
	v1 "github.com/mesatechlabs/kitten/apis/v1/todos"
)

// RouterConfig defines the configuration for the web server router
type RouterConfig struct {
	// Timeout is the maximum duration to allow the request to complete before timing out
	Timeout time.Duration

	// Cors is the CORS options to use for the web server
	Cors cors.Options
}

func InitRouter(config RouterConfig) chi.Router {
	router := chi.NewRouter()

	router.Use(
		cors.Handler(config.Cors),
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middleware.CleanPath,
		middleware.AllowContentType("application/json", "text/xml"),
		middleware.ContentCharset("UTF-8", "Latin-1"),
		middleware.Timeout(config.Timeout),
		httprate.LimitByIP(100, 1*time.Minute),
	)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, friend"))
	})

	router.Route("/api", func(r chi.Router) {
		r.Mount("/health", health.NewHealthCheckRouter())

		r.Route("/v1", func(r chi.Router) {
			r.Mount("/todos", v1.NewTodosRouter())
		})
	})

	return router
}
