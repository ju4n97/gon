package health

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mesatechlabs/kitten/tools/types"
)

type healthCheckData struct {
	AllGood bool `json:"allGood"`
}

func NewHealthCheckRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res := types.JsonResponse[healthCheckData]{
			Code:    http.StatusOK,
			Message: "OK",
			Data: healthCheckData{
				AllGood: true,
			},
		}

		json.NewEncoder(w).Encode(res)
	})

	return router
}
