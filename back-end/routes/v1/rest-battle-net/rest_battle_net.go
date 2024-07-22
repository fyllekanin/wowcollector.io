package restbattlenet

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
	"wowcollector.io/common/data"
	"wowcollector.io/entities/documents"
	"wowcollector.io/entities/response"
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
// @router /api/v1/battle-net/realms-regions [get]
func getRealmsAndRegions(w http.ResponseWriter, r *http.Request) {
	oplog := httplog.LogEntry(r.Context())
	oplog.Info("Fetching realms")

	realms, err := realmrepository.GetRepository().GetRealms()
	if err != nil {
		oplog.Error("Error fetching realms", err)
		http.Error(w, err.Error(), 500)
		return
	}

	body, err := json.Marshal(&response.RegionRealmResponse{
		Realms:  getRealms(realms),
		Regions: getRegions(),
	})
	if err != nil {
		oplog.Error("Error stringify response", err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func getRegions() []*response.RegionResponse {
	return []*response.RegionResponse{
		{Name: "EU", Value: data.REGION_EU},
		{Name: "US", Value: data.REGION_US},
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
