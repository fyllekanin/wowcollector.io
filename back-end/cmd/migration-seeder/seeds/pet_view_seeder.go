package seeds

import (
	"encoding/json"

	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
	petviewrepository "wowcollector.io/internal/repository/repositories/pet-view-repository"
	seedrepository "wowcollector.io/internal/repository/repositories/seed-repository"
)

func PetViewsSeeder() {
	var seedName = "pet-views"
	zap.L().Info("Pet views seeder started")
	result, err := seedrepository.GetRepository().IsExisting(seedName)
	if err == nil && result {
		zap.L().Info("Pet views seeder already done")
		return
	}
	byteValue, _ := GetBytesFromFile("./resources/seeds/pet-views.json")

	var petViews []documents.PetViewDocument
	err = json.Unmarshal(byteValue, &petViews)
	if err != nil {
		zap.L().Fatal("Error parsing pet-views.json")
	}

	for _, element := range petViews {
		petviewrepository.GetRepository().CreatePetView(&element)
	}
	seedrepository.GetRepository().CreateSeed(seedName)
	zap.L().Info("Pet views seeder done")
}
