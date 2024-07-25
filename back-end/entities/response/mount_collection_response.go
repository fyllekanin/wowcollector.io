package response

// Implementation struct
type MountCollectionMount struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	Id              int    `json:"id"`
	IsCollected     bool   `json:"isCollected"`
	CreatureDisplay string `json:"creatureDisplay"`
	Icon            string `json:"icon"`
}

// Original struct with recursive reference
type MountCollectionCategory struct {
	Name       string                    `json:"name"`
	Mounts     []MountCollectionMount    `json:"mounts"`
	Order      int                       `json:"order"`
	Categories []MountCollectionCategory `json:"categories"`
}

// Simplified version for Swagger documentation
type MountCollectionCategorySwagger struct {
	Name       string                 `json:"name"`
	Mounts     []MountCollectionMount `json:"mounts"`
	Order      int                    `json:"order"`
	Categories []struct{}             `json:"categories"`
}
