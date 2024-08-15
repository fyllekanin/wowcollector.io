package response

type PetAssets struct {
	LargeIcon string `json:"largeIcon"`
}

type PetCollectionPet struct {
	Name        string     `json:"name"`
	Id          int        `json:"id"`
	IsCollected bool       `json:"isCollected"`
	Assets      *PetAssets `json:"assets"`
}

type PetCollectionCategory struct {
	Name       string                  `json:"name"`
	Pets       []PetCollectionPet      `json:"pets"`
	Order      int                     `json:"order"`
	Categories []PetCollectionCategory `json:"categories"`
}

type PetCollectionCategorySwagger struct {
	Name       string             `json:"name"`
	Pets       []PetCollectionPet `json:"pets"`
	Order      int                `json:"order"`
	Categories []struct{}         `json:"categories"`
}

type PetResponse struct {
	Name           string     `json:"name"`
	Id             int        `json:"id"`
	IsUnobtainable bool       `json:"isUnobtainable"`
	Assets         *PetAssets `json:"assets"`
}
