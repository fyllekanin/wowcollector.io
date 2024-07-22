package routes

import (
	v1 "wowcollector.io/routes/v1"

	"github.com/go-chi/chi/v5"
)

func GetRoutes(r chi.Router) {
	r.Route("/api", func(r chi.Router) {
		v1.GetRoutes(r)
	})
}
