package v1

import (
	"github.com/go-chi/chi/v5"
	restbattlenet "wowcollector.io/routes/v1/rest-battle-net"
)

func GetRoutes(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		restbattlenet.GetRoutes(r)
	})
}
