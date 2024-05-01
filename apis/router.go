package apis

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

// RouterConfig defines the configuration for the web server router
type RouterConfig struct {
	// Timeout is the maximum duration to allow the request to complete before timing out
	Timeout time.Duration

	// Cors is the CORS options to use for the web server
	Cors cors.Options
}

func InitRouter(config RouterConfig) chi.Router {
	r := chi.NewRouter()

	r.Use(cors.Handler(config.Cors))
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.AllowContentType("application/json", "text/xml"))
	r.Use(middleware.Timeout(config.Timeout))
	r.Use(httprate.LimitByIP(100, 1*time.Minute))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, friend"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/example", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello, this is version 1!"))
			})
		})

		r.Route("/v2", func(r chi.Router) {
			r.Get("/example", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello, this is version 2!"))
			})
		})
	})

	return r
}
