package response

type AchievementCategoryResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	DisplayOrder int    `json:"displayOrder"`
}
