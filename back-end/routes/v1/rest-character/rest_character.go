package restcharacter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
	blizzarddata "wowcollector.io/common/data/blizzard-data"
	"wowcollector.io/entities/response"
	mountviewrepository "wowcollector.io/repository/repositories/mount-view-repository"
	"wowcollector.io/routes/v1/rest-character/aggregator"
	battlenethttp "wowcollector.io/services/http/battle-net-http"
)

func GetRoutes(r chi.Router) {
	r.Route("/character", func(r chi.Router) {
		r.Get("/{region}/{realm}/{character}", getCharacterProfile)
		r.Get("/{region}/{realm}/{character}/mounts", getCharacterMountCollection)
	})
}

// @summary Fetch character profile
// @description Get summary information about a character
// @tags Character
// @produce json
// @param region path string true "Region"
// @param realm path string true "Realm"
// @param character path string true "Character"
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

// @summary Fetch character mount collection
// @description Get mount collection for character
// @tags Character
// @produce json
// @param region path string true "Region"
// @param realm path string true "Realm"
// @param character path string true "Character"
// @success 200 {list} response.MountCollectionCategory
// @router /api/v1/character/{region}/{realm}/{character}/mounts [get]
func getCharacterMountCollection(w http.ResponseWriter, r *http.Request) {
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
	collection := battlenethttp.GetInstance().GetCharacterMountCollection(blizzarddata.FromString(region), realm, character)
	if collection == nil {
		oplog.Error("Error fetching character mount collection")
		http.Error(w, "Character mount collection not found", 400)
		return
	}
	view, _ := mountviewrepository.GetRepository().GetDefaultMountView()
	if view == nil {
		oplog.Error("Error fetching mount view")
		http.Error(w, "Mount view not found", 400)
		return
	}
	// Call aggregator
	body, err := json.Marshal(aggregator.GetMountsAggregation(
		*item,
		*collection,
		*view,
	))
	if err != nil {
		oplog.Error("Error stringify response", err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
