package restbattlenet

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	blizzarddata "wowcollector.io/internal/common/data/blizzard-data"
	errorcodes "wowcollector.io/internal/common/error-codes"
	"wowcollector.io/internal/entities/documents"
	"wowcollector.io/internal/entities/response"
	errorresponse "wowcollector.io/internal/entities/response/error"
	achievementcategoryrepository "wowcollector.io/internal/repository/repositories/achievement-category-repository"
	realmrepository "wowcollector.io/internal/repository/repositories/realm-repository"
)

func GetRoutes(r chi.Router) {
	r.Route("/battle-net", func(r chi.Router) {
		r.Get("/realms-regions", getRealmsAndRegions)
		r.Get("/root-achievement-categories", getRootAchievementCategories)
	})
}

// @summary Fetch all realms ang regions
// @description Get all the realms and regions
// @tags BattleNet
// @produce json
// @success 200 {object} response.RegionRealmResponse
// @failure 400 {object} errorresponse.ErrorResponse
// @failure 404 {object} errorresponse.ErrorResponse
// @router /api/v1/battle-net/realms-regions [get]
func getRealmsAndRegions(w http.ResponseWriter, r *http.Request) {
	zap.L().Info("Fetching collection of all realms and regions")
	realms, err := realmrepository.GetRepository().GetRealms()
	if err != nil {
		zap.L().Error("Failed fetching realms")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Failed fetching realms"))
		return
	}

	body, err := json.Marshal(&response.RegionRealmResponse{
		Realms:  getRealms(realms),
		Regions: getRegions(),
	})
	if err != nil {
		zap.L().Error("Failed to stringify response body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Invalid JSON body"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	zap.L().Info("Responded with realms and regions collection")
}

// @summary Fetch all root achievement categories
// @description Get all the achievement categories which are root (top level)
// @tags BattleNet
// @produce json
// @success 200 {object} []response.AchievementCategoryResponse
// @failure 400 {object} errorresponse.ErrorResponse
// @failure 404 {object} errorresponse.ErrorResponse
// @router /api/v1/battle-net/root-achievement-categories [get]
func getRootAchievementCategories(w http.ResponseWriter, r *http.Request) {
	zap.L().Info("Fetching achievement root categories")
	categories, err := achievementcategoryrepository.GetRepository().GetAchievementRootCategories()
	if err != nil {
		zap.L().Error("Failed fetching achievement root categories")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Failed fetching achievement root categories"))
		return
	}
	var result []*response.AchievementCategoryResponse
	for _, element := range categories {
		result = append(result, &response.AchievementCategoryResponse{
			Id:           element.Id,
			Name:         element.Name,
			DisplayOrder: element.DisplayOrder,
		})
	}

	body, err := json.Marshal(result)
	if err != nil {
		zap.L().Error("Failed to stringify response body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Invalid JSON body"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	zap.L().Info("Responded with achievement root categories")
}

func getRegions() []*response.RegionResponse {
	return []*response.RegionResponse{
		{Name: "EU", Value: blizzarddata.REGION_EU},
		{Name: "US", Value: blizzarddata.REGION_US},
	}
}

func getRealms(realms []*documents.RealmDocument) []*response.RealmResponse {
	var result []*response.RealmResponse
	for _, element := range realms {
		result = append(result, &response.RealmResponse{
			Name:   element.Name,
			Region: element.Region,
			Slug:   element.Slug,
		})
	}
	return result
}
