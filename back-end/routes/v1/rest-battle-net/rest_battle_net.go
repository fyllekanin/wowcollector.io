package restbattlenet

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	blizzarddata "wowcollector.io/common/data/blizzard-data"
	errorcodes "wowcollector.io/common/error-codes"
	"wowcollector.io/entities/documents"
	"wowcollector.io/entities/response"
	errorresponse "wowcollector.io/entities/response/error"
	realmrepository "wowcollector.io/repository/repositories/realm-repository"
)

func GetRoutes(r chi.Router) {
	r.Route("/battle-net", func(r chi.Router) {
		r.Get("/realms-regions", getRealmsAndRegions)
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
