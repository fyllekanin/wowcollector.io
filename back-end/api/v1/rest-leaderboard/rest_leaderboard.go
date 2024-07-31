package restleaderboard

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func GetRoutes(r chi.Router) {
	r.Route("/leaderboard", func(r chi.Router) {
		r.Get("/mounts", getMountsLeaderboard)
	})
}

// @summary Fetch leaderboard for mounts
// @description Get the leaderboard mounts
// @tags BattleNet
// @produce json
// @success 200 {object} response.PaginationResponse
// @failure 400 {object} errorresponse.ErrorResponse
// @failure 404 {object} errorresponse.ErrorResponse
// @router /api/v1/leaderboard/mounts [get]
func getMountsLeaderboard(w http.ResponseWriter, r *http.Request) {
	zap.L().Info("Fetching leaderboard for mounts")

	w.WriteHeader(http.StatusTeapot)
}
