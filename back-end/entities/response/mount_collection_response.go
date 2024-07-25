package response

type MountCollectionMount struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	Id              int    `json:"id"`
	IsCollected     bool   `json:"isCollected"`
	CreatureDisplay string `json:"creatureDisplay"`
	Icon            string `json:"icon"`
}

type MountCollectionCategory struct {
	Name       string                    `json:"name"`
	Mounts     []MountCollectionMount    `json:"mounts"`
	Order      int                       `json:"order"`
	Categories []MountCollectionCategory `json:"categories"`
}
