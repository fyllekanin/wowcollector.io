package restcharacter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"wowcollector.io/api/v1/rest-character/aggregator"
	blizzarddata "wowcollector.io/internal/common/data/blizzard-data"
	errorcodes "wowcollector.io/internal/common/error-codes"
	"wowcollector.io/internal/entities/response"
	errorresponse "wowcollector.io/internal/entities/response/error"
	mountviewrepository "wowcollector.io/internal/repository/repositories/mount-view-repository"
	battlenethttp "wowcollector.io/internal/services/http/battle-net-http"
)

func GetRoutes(r chi.Router) {
	r.Route("/character/{region}/{realm}/{character}", func(r chi.Router) {
		r.Get("/", getCharacterProfile)
		r.Get("/mounts", getCharacterMountCollection)
		r.Get("/achievements", getCharacterAchievementCollection)
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
// @failure 400 {object} errorresponse.ErrorResponse
// @failure 404 {object} errorresponse.ErrorResponse
// @router /api/v1/character/{region}/{realm}/{character} [get]
func getCharacterProfile(w http.ResponseWriter, r *http.Request) {
	region := chi.URLParam(r, "region")
	realm := chi.URLParam(r, "realm")
	character := chi.URLParam(r, "character")
	zap.L().Info(fmt.Sprintf("Fetching character profile for %s on realm %s in region %s", character, realm, region))

	item := battlenethttp.GetInstance().GetCharacter(blizzarddata.FromString(region), realm, character)
	if item == nil {
		zap.L().Error("Could not find the character at blizzard")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.LOADING_BATTLE_NET_DATA, fmt.Sprintf("Character %s on realm %s in region %s not found", character, realm, region)))
		return
	}

	body, err := json.Marshal(&response.CharacterProfileResponse{
		Name:    item.Name,
		Level:   item.Level,
		Realm:   item.GetRealm(),
		Faction: item.GetFaction(),
	})
	if err != nil {
		zap.L().Error("Could not stringify response body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Invalid JSON body"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	zap.L().Info("Responded with character profile")
}

// @summary Fetch character mount collection
// @description Get mount collection for character
// @tags Character
// @produce json
// @param region path string true "Region"
// @param realm path string true "Realm"
// @param character path string true "Character"
// @success 200 {object} []response.MountCollectionCategorySwagger
// @failure 400 {object} errorresponse.ErrorResponse
// @failure 404 {object} errorresponse.ErrorResponse
// @router /api/v1/character/{region}/{realm}/{character}/mounts [get]
func getCharacterMountCollection(w http.ResponseWriter, r *http.Request) {
	region := chi.URLParam(r, "region")
	realm := chi.URLParam(r, "realm")
	character := chi.URLParam(r, "character")
	zap.L().Info(fmt.Sprintf("Fetching mount collection for %s on realm %s in region %s", character, realm, region))

	item := battlenethttp.GetInstance().GetCharacter(blizzarddata.FromString(region), realm, character)
	if item == nil {
		zap.L().Error("Error finding character")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.LOADING_BATTLE_NET_DATA, fmt.Sprintf("Character %s on realm %s in region %s not found", character, realm, region)))
		return
	}
	collection := battlenethttp.GetInstance().GetCharacterMountCollection(blizzarddata.FromString(region), realm, character)
	if collection == nil {
		zap.L().Error("Error getting characters mount collection")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.LOADING_BATTLE_NET_DATA, "Character mount collection not found for character"))
		return
	}
	view, _ := mountviewrepository.GetRepository().GetDefaultMountView()
	if view == nil {
		zap.L().Error("Error finding default mount view")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.NO_DEFAULT_MOUNT_VIEW, "No default mount view available"))
		return
	}

	body, err := json.Marshal(aggregator.GetMountsAggregation(
		*item,
		*collection,
		*view,
	))
	if err != nil {
		zap.L().Error("Error stringifying response body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Invalid JSON body"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	zap.L().Info("Responded with mount collection")
}

// @summary Fetch character achievement collection
// @description Get achievement collection for character
// @tags Character
// @produce json
// @param region path string true "Region"
// @param realm path string true "Realm"
// @param character path string true "Character"
// @success 200 {object} []response.AchievementCollectionCategorySwagger
// @failure 400 {object} errorresponse.ErrorResponse
// @failure 404 {object} errorresponse.ErrorResponse
// @router /api/v1/character/{region}/{realm}/{character}/achievements [get]
func getCharacterAchievementCollection(w http.ResponseWriter, r *http.Request) {
	region := chi.URLParam(r, "region")
	realm := chi.URLParam(r, "realm")
	character := chi.URLParam(r, "character")
	zap.L().Info(fmt.Sprintf("Fetching achievement collection for %s on realm %s in region %s", character, realm, region))

	item := battlenethttp.GetInstance().GetCharacter(blizzarddata.FromString(region), realm, character)
	if item == nil {
		zap.L().Error("Error finding character")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.LOADING_BATTLE_NET_DATA, fmt.Sprintf("Character %s on realm %s in region %s not found", character, realm, region)))
		return
	}
	collection := battlenethttp.GetInstance().GetCharacterAchievementCollection(blizzarddata.FromString(region), realm, character)
	if collection == nil {
		zap.L().Error("Error getting characters achievement collection")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.LOADING_BATTLE_NET_DATA, "Character achievement collection not found for character"))
		return
	}

	body, err := json.Marshal(aggregator.GetAchievementAggregation(
		*item,
		*collection,
	))
	if err != nil {
		zap.L().Error("Error stringifying response body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Invalid JSON body"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	zap.L().Info("Responded with achievement collection")
}
