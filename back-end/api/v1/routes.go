package v1

import (
	"github.com/go-chi/chi/v5"
	restauth "wowcollector.io/api/v1/rest-auth"
	restbattlenet "wowcollector.io/api/v1/rest-battle-net"
	restcharacter "wowcollector.io/api/v1/rest-character"
	restfeedback "wowcollector.io/api/v1/rest-feedback"
	restitemview "wowcollector.io/api/v1/rest-item-view"
	restleaderboard "wowcollector.io/api/v1/rest-leaderboard"
)

func GetRoutes(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		restbattlenet.GetRoutes(r)
		restcharacter.GetRoutes(r)
		restleaderboard.GetRoutes(r)
		restfeedback.GetRoutes(r)
		restitemview.GetRoutes(r)
		restauth.GetRoutes(r)
	})
}
