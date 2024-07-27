package seeds

import (
	"encoding/json"

	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
	achievementrepository "wowcollector.io/internal/repository/repositories/achievement-repository"
	seedrepository "wowcollector.io/internal/repository/repositories/seed-repository"
)

func AchievementsSeeder() {
	var seedName = "achievements"
	zap.L().Info("Achievements seeder started")
	result, err := seedrepository.GetRepository().IsExisting(seedName)
	if err == nil && result {
		zap.L().Info("Achievements seeder already done")
		return
	}
	byteValue, _ := GetBytesFromFile("./resources/seeds/achievements.json")

	var mountViews []documents.AchievementDocument
	err = json.Unmarshal(byteValue, &mountViews)
	if err != nil {
		zap.L().Fatal("Error parsing achievements.json")
	}

	for _, element := range mountViews {
		achievementrepository.GetRepository().CreateAchievement(&element)
	}
	seedrepository.GetRepository().CreateSeed(seedName)
	zap.L().Info("Achievements seeder done")
}
