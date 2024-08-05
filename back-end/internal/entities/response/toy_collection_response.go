package response

type ToyCollectionToyAssets struct {
	LargeIcon string `json:"largeIcon"`
}

type ToyCollectionToy struct {
	Name        string                  `json:"name"`
	Id          int                     `json:"id"`
	ItemId      int                     `json:"itemId"`
	IsCollected bool                    `json:"isCollected"`
	Assets      *ToyCollectionToyAssets `json:"assets"`
}

type ToyCollectionCategory struct {
	Name       string                  `json:"name"`
	Toys       []ToyCollectionToy      `json:"toys"`
	Order      int                     `json:"order"`
	Categories []ToyCollectionCategory `json:"categories"`
}

type ToyCollectionCategorySwagger struct {
	Name       string             `json:"name"`
	Toys       []ToyCollectionToy `json:"toys"`
	Order      int                `json:"order"`
	Categories []struct{}         `json:"categories"`
}
