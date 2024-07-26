package v1

import (
	"github.com/go-chi/chi/v5"
	restbattlenet "wowcollector.io/api/v1/rest-battle-net"
	restcharacter "wowcollector.io/api/v1/rest-character"
)

func GetRoutes(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		restbattlenet.GetRoutes(r)
		restcharacter.GetRoutes(r)
	})
}
