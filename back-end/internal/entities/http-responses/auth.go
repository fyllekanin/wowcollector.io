package httpresponses

type BattleNetAuth struct {
	AccessToken string `json:"access_token"`
}

type BattleNetUserInfo struct {
	BattleTag string `json:"battletag"`
}
