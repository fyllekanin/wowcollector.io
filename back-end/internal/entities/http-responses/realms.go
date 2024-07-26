package httpresponses

type BattleNetRealm struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type BattleNetRealms struct {
	Realms []BattleNetRealm
}
