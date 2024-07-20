package v1

import (
	"github.com/go-chi/chi/v5"
	helloworld "wowcollector.io/routes/v1/hello-world"
)

func GetRoutes(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		helloworld.GetRoutes(r)
	})
}
