package seeds

import (
	"encoding/json"

	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
	achievementcategoryrepository "wowcollector.io/internal/repository/repositories/achievement-category-repository"
	seedrepository "wowcollector.io/internal/repository/repositories/seed-repository"
)

func AchievementCategoriesSeeder() {
	var seedName = "achievement-categories"
	zap.L().Info("Achievement categories seeder started")
	result, err := seedrepository.GetRepository().IsExisting(seedName)
	if err == nil && result {
		zap.L().Info("Achievement categories seeder already done")
		return
	}
	byteValue, _ := GetBytesFromFile("./resources/seeds/achievement-categories.json")

	var mountViews []documents.AchievementCategoryDocument
	err = json.Unmarshal(byteValue, &mountViews)
	if err != nil {
		zap.L().Fatal("Error parsing achievement-categories.json")
	}

	for _, element := range mountViews {
		achievementcategoryrepository.GetRepository().CreateAchievementCategory(&element)
	}
	seedrepository.GetRepository().CreateSeed(seedName)
	zap.L().Info("Achievement categories seeder done")
}
