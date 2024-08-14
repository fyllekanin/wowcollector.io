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
	mountrepository "wowcollector.io/internal/repository/repositories/mount-repository"
	realmrepository "wowcollector.io/internal/repository/repositories/realm-repository"
)

func GetRoutes(r chi.Router) {
	r.Route("/battle-net", func(r chi.Router) {
		r.Get("/realms-regions", getRealmsAndRegions)
		r.Get("/mounts", getMounts)
	})
}

// @summary Fetch all mounts
// @description Get all the realms and regions
// @tags BattleNet
// @produce json
// @success 200 {object} []response.MountResponse
// @failure 400 {object} errorresponse.ErrorResponse
// @router /api/v1/battle-net/mounts [get]
func getMounts(w http.ResponseWriter, r *http.Request) {
	zap.L().Info("Fetching collection of all mounts")
	mounts, err := mountrepository.GetRepository().GetMounts()
	if err != nil {
		zap.L().Error("Failed fetching mounts")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorresponse.GenerateErrorBody(errorcodes.INTERNAL_ERROR, "Failed fetching mounts"))
		return
	}

	var result []*response.MountResponse
	for _, element := range mounts {
		result = append(result, &response.MountResponse{
			Name:           element.Name,
			Description:    element.Description,
			Id:             element.Id,
			IsUnobtainable: element.IsUnobtainable,
			Assets: &response.MountAssets{
				Display:   element.Assets.Display,
				SmallIcon: element.Assets.SmallIcon,
				LargeIcon: element.Assets.LargeIcon,
			},
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
	zap.L().Info("Responded with mounts")
}

// @summary Fetch all realms and regions
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
