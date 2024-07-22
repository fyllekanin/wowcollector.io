package restcharacter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
	blizzarddata "wowcollector.io/common/data/blizzard-data"
	"wowcollector.io/entities/response"
	battlenethttp "wowcollector.io/services/http/battle-net-http"
)

func GetRoutes(r chi.Router) {
	r.Route("/character", func(r chi.Router) {
		r.Get("/{region}/{realm}/{character}", getCharacterProfile)
	})
}

// @summary Fetch all realms ang regions
// @description Get all the realms and regions
// @tags BattleNet
// @produce json
// @success 200 {object} response.CharacterProfileResponse
// @router /api/v1/character/{region}/{realm}/{character} [get]
func getCharacterProfile(w http.ResponseWriter, r *http.Request) {
	oplog := httplog.LogEntry(r.Context())
	region := chi.URLParam(r, "region")
	realm := chi.URLParam(r, "realm")
	character := chi.URLParam(r, "character")
	oplog.Info(fmt.Sprintf("Fetching character %s in region %s on realm %s", character, region, realm))

	item := battlenethttp.GetInstance().GetCharacter(blizzarddata.FromString(region), realm, character)
	if item == nil {
		oplog.Error("Error fetching character")
		http.Error(w, "Character not found", 404)
		return
	}

	body, err := json.Marshal(&response.CharacterProfileResponse{
		Name:    item.Name,
		Level:   item.Level,
		Realm:   item.GetRealm(),
		Faction: item.GetFaction(),
	})
	if err != nil {
		oplog.Error("Error stringify response", err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
