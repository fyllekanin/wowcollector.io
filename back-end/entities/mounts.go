package battlenetEntities

type BattleNetMountIndex struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type BattleNetMountsIndex struct {
	Mounts []BattleNetMountIndex `json:"mounts"`
}
