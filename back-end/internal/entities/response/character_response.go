package response

type CharacterProfileAssets struct {
	Avatar  string `json:"avatar"`
	Inset   string `json:"inset"`
	MainRaw string `json:"mainRaw"`
}

type CharacterProfileResponse struct {
	Name    string                  `json:"name"`
	Level   int                     `json:"level"`
	Realm   string                  `json:"realm"`
	Faction string                  `json:"faction"`
	Assets  *CharacterProfileAssets `json:"assets"`
}
