package httpresponses

type BattleNetAuth struct {
	AccessCode string `json:"access_code"`
}

type BattleNetUserInfo struct {
	BattleTag string `json:"battletag"`
}
