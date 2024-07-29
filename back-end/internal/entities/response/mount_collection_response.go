package response

type MountCollectionMountAssets struct {
	Display   string `json:"display"`
	SmallIcon string `json:"smallIcon"`
	LargeIcon string `json:"largeIcon"`
}

type MountCollectionMount struct {
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	Id          int                         `json:"id"`
	IsCollected bool                        `json:"isCollected"`
	Assets      *MountCollectionMountAssets `json:"assets"`
}

type MountCollectionCategory struct {
	Name       string                    `json:"name"`
	Mounts     []MountCollectionMount    `json:"mounts"`
	Order      int                       `json:"order"`
	Categories []MountCollectionCategory `json:"categories"`
}

type MountCollectionCategorySwagger struct {
	Name       string                 `json:"name"`
	Mounts     []MountCollectionMount `json:"mounts"`
	Order      int                    `json:"order"`
	Categories []struct{}             `json:"categories"`
}
