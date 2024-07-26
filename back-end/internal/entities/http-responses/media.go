package httpresponses

type BattleNetAsset struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type BattleNetMedia struct {
	Assets []BattleNetAsset `json:"assets"`
	Id     int              `json:"id"`
}
