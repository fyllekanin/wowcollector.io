package httpresponses

type BattleNetMountIndex struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type BattleNetMountsIndex struct {
	Mounts []BattleNetMountIndex `json:"mounts"`
}

type BattleNetMountSource struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type BattleNetMount struct {
	Id                         int                       `json:"id"`
	Name                       string                    `json:"name"`
	Description                string                    `json:"description"`
	Source                     *BattleNetMountSource     `json:"source"`
	Faction                    *BattleNetFaction         `json:"faction"`
	CreatureDisplays           []*BattleNetEntityDisplay `json:"creature_displays"`
	ShouldExcludeIfUncollected bool                      `json:"should_exclude_if_uncollected"`
}

type BattleNetCharacterMountDetails struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type BattleNetCharacterMount struct {
	Mount *BattleNetCharacterMountDetails `json:"mount"`
}

type BattleNetCharacterMountCollection struct {
	Mounts []*BattleNetCharacterMount `json:"mounts"`
}
