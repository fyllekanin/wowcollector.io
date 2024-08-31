package response

type RealmResponse struct {
	Name   string `json:"name"`
	Region string `json:"region"`
	Slug   string `json:"slug"`
}

type RegionResponse struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type RegionRealmResponse struct {
	Realms  []*RealmResponse  `json:"realms"`
	Regions []*RegionResponse `json:"regions"`
}
