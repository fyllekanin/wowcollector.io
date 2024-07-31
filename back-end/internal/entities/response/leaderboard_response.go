package response

type LeaderboardEntry struct {
	Character string `json:"character"`
	Realm     string `json:"realm"`
	Region    string `json:"region"`
	Count     int    `json:"count"`
}

type PaginationResponse struct {
	Total int                 `json:"total"`
	Items []*LeaderboardEntry `json:"items"`
}
