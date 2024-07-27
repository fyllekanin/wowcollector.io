package httpresponses

type BattleNetAchievementCategory struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type BattleNetAchievementCategoryIndex struct {
	Categories     []*BattleNetAchievementCategory `json:"categories"`
	RootCategories []*BattleNetAchievementCategory `json:"root_categories"`
}

type BattleNetAchievementIndexItem struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type BattleNetAchievementIndex struct {
	Achievements   []*BattleNetAchievementIndexItem `json:"achievements"`
	ParentCategory *BattleNetAchievementCategory    `json:"parent_category"`
	DisplayOrder   int                              `json:"display_order"`
}

type BattleNetAchievementMedia struct {
	Key BattleNetHref `json:"key"`
}

type BattleNetAchievement struct {
	Id            int                       `json:"id"`
	Name          string                    `json:"name"`
	Description   string                    `json:"description"`
	Points        int                       `json:"points"`
	IsAccountWide bool                      `json:"is_account_wide"`
	Media         BattleNetAchievementMedia `json:"media"`
	DisplayOrder  int                       `json:"display_order"`
}
