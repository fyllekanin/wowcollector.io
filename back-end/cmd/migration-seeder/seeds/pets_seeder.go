package seeds

import (
	"encoding/json"

	"go.uber.org/zap"
	"wowcollector.io/internal/entities/documents"
	petrepository "wowcollector.io/internal/repository/repositories/pet-repository"
	seedrepository "wowcollector.io/internal/repository/repositories/seed-repository"
)

func PetsSeeder() {
	var seedName = "pets"
	zap.L().Info("Pets seeder started")
	result, err := seedrepository.GetRepository().IsExisting(seedName)
	if err == nil && result {
		zap.L().Info("Pets seeder already done")
		return
	}
	byteValue, _ := GetBytesFromFile("./resources/seeds/pets.json")

	var pets []documents.PetDocument
	err = json.Unmarshal(byteValue, &pets)
	if err != nil {
		zap.L().Fatal("Error parsing pets.json")
	}

	for _, element := range pets {
		petrepository.GetRepository().CreatePet(&element)
	}
	seedrepository.GetRepository().CreateSeed(seedName)
	zap.L().Info("Pets seeder done")
}
