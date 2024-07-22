package battlenetEntities

type BattleNetFaction struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type BattleNetCharacter struct {
	Faction BattleNetFaction `json:"faction"`
	Realm   BattleNetRealm   `json:"realm"`
	Level   int              `json:"level"`
	Name    string           `json:"name"`
}
