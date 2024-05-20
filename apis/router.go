package apis

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/go-chi/render"
	"github.com/jm2097/gon/apis/health"
	todosV1 "github.com/jm2097/gon/apis/v1/todos"
)

// RouterConfig defines the configuration for the web server router.
type RouterConfig struct {
	// Timeout is the maximum duration to allow the request to complete before timing out
	Timeout time.Duration

	// Cors is the CORS options to use for the web server
	Cors cors.Options

	// RateLimit is the rate limit configuration for the web server
	RateLimit RouterConfigRateLimit
}

type RouterConfigRateLimit struct {
	RequestsLimit  int
	RequestsWindow time.Duration
}

func InitRouter(config *RouterConfig) chi.Router {
	router := chi.NewRouter()

	router.Use(
		cors.Handler(config.Cors),
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middleware.CleanPath,
		middleware.Timeout(config.Timeout),
		render.SetContentType(render.ContentTypeJSON),
		httprate.LimitByIP(config.RateLimit.RequestsLimit, config.RateLimit.RequestsWindow),
	)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, friend"))
	})

	router.Route("/api", func(r chi.Router) {
		r.Mount("/health", health.NewHealthCheckRouter())

		// API version 1
		r.Route("/v1", func(r chi.Router) {
			r.Mount("/todos", todosV1.NewTodosRouter())
		})
	})

	return router
}
