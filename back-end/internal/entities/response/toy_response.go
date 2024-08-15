package response

type ToyAssets struct {
	LargeIcon string `json:"largeIcon"`
}

type ToyCollectionToy struct {
	Name        string     `json:"name"`
	Id          int        `json:"id"`
	ItemId      int        `json:"itemId"`
	IsCollected bool       `json:"isCollected"`
	Assets      *ToyAssets `json:"assets"`
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

type ToyResponse struct {
	Name           string     `json:"name"`
	Id             int        `json:"id"`
	ItemId         int        `json:"itemId"`
	Assets         *ToyAssets `json:"assets"`
	IsUnobtainable bool       `json:"isUnobtainable"`
}
