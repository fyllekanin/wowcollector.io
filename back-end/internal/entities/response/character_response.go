package response

type CharacterProfileResponse struct {
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Realm   string `json:"realm"`
	Faction string `json:"faction"`
}
