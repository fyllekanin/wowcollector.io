package response

import "wowcollector.io/common/data"

type RealmResponse struct {
	Name   string               `json:"name"`
	Region data.BattleNetRegion `json:"region"`
	Slug   string               `json:"slug"`
}

type RegionResponse struct {
	Name  string               `json:"name"`
	Value data.BattleNetRegion `json:"value"`
}

type RegionRealmResponse struct {
	Realms  []*RealmResponse  `json:"realms"`
	Regions []*RegionResponse `json:"regions"`
}
