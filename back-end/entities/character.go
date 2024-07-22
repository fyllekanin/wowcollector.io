package battlenetEntities

type BattleNetFaction struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type BattleNetCharacter struct {
	Faction *BattleNetFaction `json:"faction"`
	Realm   *BattleNetRealm   `json:"realm"`
	Level   int               `json:"level"`
	Name    string            `json:"name"`
}

func (b *BattleNetCharacter) GetFaction() string {
	if b.Faction != nil {
		return b.Faction.Type
	}
	return ""
}

func (b *BattleNetCharacter) GetRealm() string {
	if b.Realm != nil {
		return b.Realm.Slug
	}
	return ""
}
