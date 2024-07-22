package battlenetEntities

type BattleNetRealm struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type BattleNetRealms struct {
	Realms []BattleNetRealm
}
