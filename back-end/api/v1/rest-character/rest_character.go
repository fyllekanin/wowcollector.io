package restcharacter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	achievementaggregator "wowcollector.io/api/v1/rest-character/aggregators/achievement-aggregator"
	mountaggregator "wowcollector.io/api/v1/rest-character/aggregators/mount-aggregator"
	petaggregator "wowcollector.io/api/v1/rest-character/aggregators/pet-aggregator"
	toyaggregator "wowcollector.io/api/v1/rest-character/aggregators/toy-aggregator"
	errorcodes "wowcollector.io/internal/common/error-codes"
	"wowcollector.io/internal/entities/documents"
	httpresponses "wowcollector.io/internal/entities/http-responses"
	"wowcollector.io/internal/entities/response"
	errorresponse "wowcollector.io/internal/entities/response/error"
	mountleaderboardrepository "wowcollector.io/internal/repository/repositories/mount-leaderboard-repository"
	mountviewrepository "wowcollector.io/internal/repository/repositories/mount-view-repository"
	petviewrepository "wowcollector.io/internal/repository/repositories/pet-view-repository"
	toyviewrepository "wowcollector.io/internal/repository/repositories/toy-view-repository"
	battlenethttp "wowcollector.io/internal/services/http/battle-net-http"
)

func GetRoutes(r chi.Router) {
	r.Route("/character/{region}/{realm}/{character}", func(r chi.Router) {
		r.Get("/", getCharacterProfile)
		r.Get("/mounts", getCharacterMountCollection)
		r.Get("/achievements", getCharacterAchievementCollection)
		r.Get("/toys", getCharacterToyCollection)
		r.Get("/pets", getCharacterPetCollection)
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

	item := battlenethttp.GetInstance().GetCharacter(region, realm, character)
	if item == nil {
		zap.L().Error("Could not find the character at blizzard")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.LOADING_BATTLE_NET_DATA, fmt.Sprintf("Character %s on realm %s in region %s not found", character, realm, region)))
		return
	}

	media := battlenethttp.GetInstance().GetCharacterMedia(region, realm, character)
	body, err := json.Marshal(&response.CharacterProfileResponse{
		Name:    item.Name,
		Level:   item.Level,
		Realm:   item.GetRealm(),
		Faction: item.GetFaction(),
		Assets:  getCharacterAssets(media),
		Region:  region,
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

func getCharacterAssets(media *httpresponses.BattleNetMedia) *response.CharacterProfileAssets {
	if media == nil {
		return nil
	}
	return &response.CharacterProfileAssets{
		Avatar:  media.GetAssetByKey("avatar"),
		Inset:   media.GetAssetByKey("inset"),
		MainRaw: media.GetAssetByKey("main-raw"),
	}
}

// @summary Fetch character mount collection
// @description Get mount collection for character
// @tags Character
// @produce json
// @param region path string true "Region"
// @param realm path string true "Realm"
// @param character path string true "Character"
// @param viewId query string false "ViewID"
// @success 200 {object} []response.MountCollectionCategorySwagger
// @failure 400 {object} errorresponse.ErrorResponse
// @failure 404 {object} errorresponse.ErrorResponse
// @router /api/v1/character/{region}/{realm}/{character}/mounts [get]
func getCharacterMountCollection(w http.ResponseWriter, r *http.Request) {
	region := chi.URLParam(r, "region")
	realm := chi.URLParam(r, "realm")
	character := chi.URLParam(r, "character")
	viewId := r.URL.Query().Get("viewId")
	zap.L().Info(fmt.Sprintf("Fetching mount collection for %s on realm %s in region %s", character, realm, region))

	item := battlenethttp.GetInstance().GetCharacter(region, realm, character)
	if item == nil {
		zap.L().Error("Error finding character")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.LOADING_BATTLE_NET_DATA, fmt.Sprintf("Character %s on realm %s in region %s not found", character, realm, region)))
		return
	}
	collection := battlenethttp.GetInstance().GetCharacterMountCollection(region, realm, character)
	if collection == nil {
		zap.L().Error("Error getting characters mount collection")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.LOADING_BATTLE_NET_DATA, "Character mount collection not found for character"))
		return
	}
	go updateMountLeaderBoard(character, realm, region, len(collection.Mounts))

	var view *documents.MountViewDocument
	if viewId == "" {
		viewResult, err := mountviewrepository.GetRepository().GetDefaultMountView()
		if err != nil {
			zap.L().Error("Error finding default mount view")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorresponse.GenerateErrorBody(errorcodes.NO_DEFAULT_MOUNT_VIEW, "No default mount view available"))
			return
		}
		view = viewResult
	} else {
		viewResult, err := mountviewrepository.GetRepository().GetMountView(viewId)
		if err != nil {
			zap.L().Error("Error finding mount view")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorresponse.GenerateErrorBody(errorcodes.NO_MOUNT_VIEW_WITH_NAME, "No mount view with id: "+viewId))
			return
		}
		view = viewResult
	}

	body, err := json.Marshal(mountaggregator.GetMountsAggregation(
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

// @summary Fetch character toy collection
// @description Get toy collection for character
// @tags Character
// @produce json
// @param region path string true "Region"
// @param realm path string true "Realm"
// @param character path string true "Character"
// @param viewId query string false "ViewID"
// @success 200 {object} []response.ToyCollectionCategorySwagger
// @failure 400 {object} errorresponse.ErrorResponse
// @failure 404 {object} errorresponse.ErrorResponse
// @router /api/v1/character/{region}/{realm}/{character}/toys [get]
func getCharacterToyCollection(w http.ResponseWriter, r *http.Request) {
	region := chi.URLParam(r, "region")
	realm := chi.URLParam(r, "realm")
	character := chi.URLParam(r, "character")
	viewId := r.URL.Query().Get("viewId")
	zap.L().Info(fmt.Sprintf("Fetching toy collection for %s on realm %s in region %s", character, realm, region))

	item := battlenethttp.GetInstance().GetCharacter(region, realm, character)
	if item == nil {
		zap.L().Error("Error finding character")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.LOADING_BATTLE_NET_DATA, fmt.Sprintf("Character %s on realm %s in region %s not found", character, realm, region)))
		return
	}
	collection := battlenethttp.GetInstance().GetCharacterToyCollection(region, realm, character)
	if collection == nil {
		zap.L().Error("Error getting characters toy collection")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.LOADING_BATTLE_NET_DATA, "Character toy collection not found for character"))
		return
	}

	var view *documents.ToyViewDocument
	if viewId == "" {
		viewResult, err := toyviewrepository.GetRepository().GetDefaultToyView()
		if err != nil {
			zap.L().Error("Error finding default toy view")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorresponse.GenerateErrorBody(errorcodes.NO_DEFAULT_TOY_VIEW, "No default toy view available"))
			return
		}
		view = viewResult
	} else {
		viewResult, err := toyviewrepository.GetRepository().GetToyView(viewId)
		if err != nil {
			zap.L().Error("Error finding toy view")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorresponse.GenerateErrorBody(errorcodes.NO_TOY_VIEW_WITH_NAME, "No toy view with id: "+viewId))
			return
		}
		view = viewResult
	}

	body, err := json.Marshal(toyaggregator.GetToysAggregation(
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
	zap.L().Info("Responded with toy collection")
}

// @summary Fetch character pet collection
// @description Get pet collection for character
// @tags Character
// @produce json
// @param region path string true "Region"
// @param realm path string true "Realm"
// @param character path string true "Character"
// @param viewId query string false "ViewID"
// @success 200 {object} []response.PetCollectionCategorySwagger
// @failure 400 {object} errorresponse.ErrorResponse
// @failure 404 {object} errorresponse.ErrorResponse
// @router /api/v1/character/{region}/{realm}/{character}/pets [get]
func getCharacterPetCollection(w http.ResponseWriter, r *http.Request) {
	region := chi.URLParam(r, "region")
	realm := chi.URLParam(r, "realm")
	character := chi.URLParam(r, "character")
	viewId := r.URL.Query().Get("viewId")
	zap.L().Info(fmt.Sprintf("Fetching pet collection for %s on realm %s in region %s", character, realm, region))

	item := battlenethttp.GetInstance().GetCharacter(region, realm, character)
	if item == nil {
		zap.L().Error("Error finding character")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.LOADING_BATTLE_NET_DATA, fmt.Sprintf("Character %s on realm %s in region %s not found", character, realm, region)))
		return
	}
	collection := battlenethttp.GetInstance().GetCharacterPetCollection(region, realm, character)
	if collection == nil {
		zap.L().Error("Error getting characters pet collection")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.LOADING_BATTLE_NET_DATA, "Character pet collection not found for character"))
		return
	}

	var view *documents.PetViewDocument
	if viewId == "" {
		viewResult, err := petviewrepository.GetRepository().GetDefaultPetView()
		if err != nil {
			zap.L().Error("Error finding default pet view")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorresponse.GenerateErrorBody(errorcodes.NO_PET_VIEW_WITH_NAME, "No default pet view available"))
			return
		}
		view = viewResult
	} else {
		viewResult, err := petviewrepository.GetRepository().GetPetView(viewId)
		if err != nil {
			zap.L().Error("Error finding pet view")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorresponse.GenerateErrorBody(errorcodes.NO_PET_VIEW_WITH_NAME, "No pet view with id: "+viewId))
			return
		}
		view = viewResult
	}

	body, err := json.Marshal(petaggregator.GetPetsAggregation(
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
	zap.L().Info("Responded with pet collection")
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

	item := battlenethttp.GetInstance().GetCharacter(region, realm, character)
	if item == nil {
		zap.L().Error("Error finding character")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.LOADING_BATTLE_NET_DATA, fmt.Sprintf("Character %s on realm %s in region %s not found", character, realm, region)))
		return
	}
	collection := battlenethttp.GetInstance().GetCharacterAchievementCollection(region, realm, character)
	if collection == nil {
		zap.L().Error("Error getting characters achievement collection")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.LOADING_BATTLE_NET_DATA, "Character achievement collection not found for character"))
		return
	}

	categories := achievementaggregator.GetAchievementAggregation(
		*item,
		*collection,
	)
	body, err := json.Marshal(categories)
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

func updateMountLeaderBoard(character string, realm string, region string, count int) {
	existing, _ := mountleaderboardrepository.GetRepository().GetLeaderBoardEntry(character, realm, region)

	document := &documents.LeaderboardDocument{
		Character: character,
		Realm:     realm,
		Region:    region,
		Count:     count,
	}

	if existing == nil {
		document.ObjectID = primitive.NewObjectID()
		mountleaderboardrepository.GetRepository().CreateLeaderboardEntry(document)
		zap.L().Info("Added new mount leader board entry")
	} else {
		document.ObjectID = existing.ObjectID
		if existing.Count != count {
			mountleaderboardrepository.GetRepository().UpdateLeaderboardEntry(document)
			zap.L().Info("Updated mount leader board entry")
		}
	}
}
