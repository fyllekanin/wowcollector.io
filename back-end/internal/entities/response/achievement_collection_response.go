package response

type AchievementCollectionAchievement struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Points       int    `json:"points"`
	DisplayOrder int    `json:"displayOrder"`
	Icon         string `json:"icon"`
	IsCompleted  bool   `json:"isCompleted"`
}

type AchievementCollectionCategory struct {
	Name         string                             `json:"name"`
	Achievements []AchievementCollectionAchievement `json:"achievements"`
	Categories   []AchievementCollectionCategory    `json:"categories"`
	DisplayOrder int                                `json:"displayOrder"`
}

type AchievementCollectionCategorySwagger struct {
	Name         string                             `json:"name"`
	Achievements []AchievementCollectionAchievement `json:"achievements"`
	Categories   []struct{}                         `json:"categories"`
	DisplayOrder int                                `json:"displayOrder"`
}
