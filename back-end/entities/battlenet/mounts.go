package battlenetEntities

type BattleNetMountIndex struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type BatleNetMountsIndex struct {
	Mounts []BattleNetMountIndex `json:"mounts"`
}
