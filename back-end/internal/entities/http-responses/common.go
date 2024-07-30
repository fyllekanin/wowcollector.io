package httpresponses

type BattleNetSource struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type BattleNetItemProperty struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type BattleNetMediaProperty struct {
	Id int `json:"id"`
}
