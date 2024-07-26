package tasks

import (
	"fmt"
	"slices"

	"go.uber.org/zap"
	blizzarddata "wowcollector.io/internal/common/data/blizzard-data"
	httpresponses "wowcollector.io/internal/entities/http-responses"
	battlenethttp "wowcollector.io/internal/services/http/battle-net-http"
)

func ScanAchievements(region blizzarddata.BattleNetRegion) {
	zap.L().Info(fmt.Sprintf("Starting scan of achievements for region %s\n", region))
	index := battlenethttp.GetInstance().GetAchievementCategoryIndex(region)
	allCategories := slices.Concat(index.Categories, index.RootCategories)

	for _, category := range allCategories {
		go runCategory(region, category)
	}
}

func runCategory(region blizzarddata.BattleNetRegion, category httpresponses.BattleNetAchievementCategory) {
	zap.L().Info(fmt.Sprintf("Scanning achievement category %d", category.Id))
	achievementIndex := battlenethttp.GetInstance().GetAchievementCategory(region, category.Id)
	if achievementIndex == nil {
		zap.L().Info(fmt.Sprintf("Could not get achievement category %d", category.Id))
		return
	}

	// Add achievement category in database

	// Loop achievements and do the same
}
