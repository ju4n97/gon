package health

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHealthCheckRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
	})

	return router
}
