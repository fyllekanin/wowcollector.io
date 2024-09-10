package httpresponses

type BattleNetAuth struct {
	AccessToken string `json:"access_token"`
}

type BattleNetUserInfo struct {
	BattleTag string `json:"battletag"`
}

type DiscordAuth struct {
	AccessToken string `json:"access_token"`
}

type DiscordUserInfo struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Bot           bool   `json:"bot"`
}
