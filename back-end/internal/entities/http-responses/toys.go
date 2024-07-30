package httpresponses

type BattleNetToysIndexToy struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type BattleNetToysIndex struct {
	Toys []*BattleNetToysIndexToy `json:"toys"`
}

type BattleNetToy struct {
	Id                         int                     `json:"id"`
	Item                       *BattleNetItemProperty  `json:"item"`
	Source                     *BattleNetSource        `json:"source"`
	ShouldExcludeIfUncollected bool                    `json:"should_exclude_if_uncollected"`
	Media                      *BattleNetMediaProperty `json:"media"`
}
