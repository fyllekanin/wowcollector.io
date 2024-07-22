package response

import "wowcollector.io/common/data"

type RealmResponse struct {
	Name   string
	Region data.BattleNetRegion
	Slug   string
}

type RegionResponse struct {
	Name  string
	Value data.BattleNetRegion
}

type RegionRealmResponse struct {
	Realms  []*RealmResponse
	Regions []*RegionResponse
}
