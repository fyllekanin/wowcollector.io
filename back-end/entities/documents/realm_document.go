package documents

import "wowcollector.io/common/data"

type RealmDocument struct {
	Id     int64
	Name   string
	Slug   string
	Region data.BattleNetRegion
}
