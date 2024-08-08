package httpresponses

type BattleNetPetsIndexPet struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type BattleNetPetsIndex struct {
	Pets []*BattleNetPetsIndexPet `json:"pets"`
}

type BattleNetPetCreature struct {
	Id int `json:"id"`
}

type BattleNetPet struct {
	Id                         int                   `json:"id"`
	Name                       string                `json:"name"`
	Source                     *BattleNetSource      `json:"source"`
	ShouldExcludeIfUncollected bool                  `json:"should_exclude_if_uncollected"`
	Icon                       string                `json:"icon"`
	Creature                   *BattleNetPetCreature `json:"creature"`
}

type BattleNetCharacterPet struct {
	Id int `json:"id"`
}

type BattleNetCharacterPetCollection struct {
	Pets []*BattleNetCharacterPet `json:"pets"`
}
