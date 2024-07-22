package response

import (
	blizzarddata "wowcollector.io/common/data/blizzard-data"
)

type RealmResponse struct {
	Name   string                       `json:"name"`
	Region blizzarddata.BattleNetRegion `json:"region"`
	Slug   string                       `json:"slug"`
}

type RegionResponse struct {
	Name  string                       `json:"name"`
	Value blizzarddata.BattleNetRegion `json:"value"`
}

type RegionRealmResponse struct {
	Realms  []*RealmResponse  `json:"realms"`
	Regions []*RegionResponse `json:"regions"`
}
